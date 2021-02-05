package main

import (
	"fmt"
	"go_boilerplate/api"
	"go_boilerplate/config"
	"go_boilerplate/pkg/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.App, cfg.LogLevel)

	psqlConnString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	db, err := sqlx.Connect("postgres", psqlConnString)

	if err != nil {
		log.Panic("error connecting to postgres", logger.Error(err))
	}

	apiServer := api.New(cfg, log, db)

	apiServer.Run(cfg.HTTPPort)
}
