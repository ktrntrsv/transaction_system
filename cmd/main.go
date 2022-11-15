package main

import (
	"context"
	"github.com/ktrntrsv/transactionService/internal/adapters"
	"github.com/ktrntrsv/transactionService/internal/adapters/db"
	"github.com/ktrntrsv/transactionService/internal/config"
	"github.com/ktrntrsv/transactionService/internal/domain"
	"github.com/ktrntrsv/transactionService/pkg/logger"
	"github.com/ktrntrsv/transactionService/pkg/postgresql"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig("./config.yml")
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	l := logger.New(cfg.Logger.Level)

	postgresSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Postgres)
	if err != nil {
		l.Fatal(err)
	}
	l.Info("connected to postgreSQL")

	dbClient := db.Database{Pool: postgresSQLClient}
	accRepository := db.NewAccountRepository(&dbClient, l)

	accUsecase := domain.NewAccountUsecase(accRepository)

	e := adapters.SetRoutes(l, accUsecase)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
