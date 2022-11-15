package adapters

import (
	"errors"
	"github.com/google/uuid"
	"github.com/ktrntrsv/transactionService/internal/domain"
	"github.com/ktrntrsv/transactionService/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

// Handler
func transferHandler(c echo.Context, accUc *domain.AccountUsecase) error {
	type transactionRequest struct {
		AccountFromId uuid.UUID `json:"accountFromId"`
		AccountToId   uuid.UUID `json:"accountToId"`
		Amount        float64   `json:"amount"`
	}
	var reqBody transactionRequest

	// check terms
	err := c.Bind(&reqBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "wrong request parameters")
	}

	if reqBody.AccountToId == reqBody.AccountFromId {
		return c.JSON(http.StatusBadRequest, "account from id = account to id")
	}

	// transferring
	err = accUc.TransferMoney(c.Request().Context(), reqBody.AccountFromId, reqBody.AccountToId, reqBody.Amount)
	if err != nil {
		if errors.Is(err, domain.ErrorNotEnoughFunds) {
			return c.JSON(http.StatusBadRequest, "not enough funds")
		}
		if errors.Is(err, domain.ErrAccountNotFound) {
			return c.JSON(http.StatusBadRequest, "account not found")
		}
		return c.JSON(http.StatusInternalServerError, "can not implement transaction")
	}

	return c.JSON(http.StatusOK, "All done")
}

func getBalance(c echo.Context, accUc *domain.AccountUsecase, userID uuid.UUID) error {
	type getBalanceResponse struct {
		Balance float64 `json:"balance"`
	}

	balance, err := accUc.GetBalance(c.Request().Context(), userID)
	if err != nil {
		if errors.Is(err, domain.ErrAccountNotFound) {
			return c.JSON(http.StatusBadRequest, "account not found")
		}
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}
	resp := getBalanceResponse{balance}
	return c.JSON(http.StatusOK, resp)
}

func SetRoutes(l logger.Interface, accUc *domain.AccountUsecase) *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/transfer", func(c echo.Context) error {
		l.Info("In /transfer")
		return transferHandler(c, accUc)
	})

	e.GET("/balance/:id", func(c echo.Context) error {
		accId := c.Param("id")
		l.Info("In get balance, " + accId)
		id, err := uuid.Parse(accId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "wrong uuid param")
		}
		return getBalance(c, accUc, id)
	})

	e.GET("/healthz", func(c echo.Context) error {
		l.Info("Health")
		return c.JSON(http.StatusOK, "Health")
	})

	return e

}
