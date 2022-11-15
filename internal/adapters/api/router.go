package api

import (
	"errors"
	uuid2 "github.com/gofrs/uuid"
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
		return c.String(http.StatusBadRequest, "wrong request parameters")
	}

	if reqBody.AccountToId == reqBody.AccountFromId {
		return c.String(http.StatusBadRequest, "account from id = account to id")
	}

	// transferring
	err = accUc.TransferMoney(c.Request().Context(), uuid.UUID(uuid2.UUID(reqBody.AccountFromId)), uuid.UUID(uuid2.UUID(reqBody.AccountToId)), reqBody.Amount)
	if err != nil {
		if errors.Is(err, domain.ErrorNotEnoughFunds) {
			return c.String(http.StatusBadRequest, "not enough funds")
		}
		if errors.Is(err, domain.ErrAccountNotFound) {
			return c.String(http.StatusBadRequest, "account not found")
		}
		return c.String(http.StatusInternalServerError, "can not implement transaction")
	}

	return c.String(http.StatusOK, "All done")
}

//func getBalance

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

	//e.GET("/balance/:id", func(c echo.Context) error {
	//	userId := c.Param("id")
	//	l.Info("In get balance, " + userId)
	//	return getBalance(c, accUc, userId)
	//})

	e.GET("/healthz", func(c echo.Context) error {
		l.Info("Health")
		return c.String(http.StatusOK, "Health")
	})

	return e

}
