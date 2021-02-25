package application

import (
	"go_boilerplate/config"
	"go_boilerplate/pkg/logger"
	"go_boilerplate/storage"

	"github.com/jmoiron/sqlx"
)

// Application ...
type Application struct {
	cfg             config.Config
	log             logger.Logger
	storagePostgres storage.PostgresStorageI
}

// New ...
func New(cfg config.Config, log logger.Logger, db *sqlx.DB) *Application {
	return &Application{
		cfg:             cfg,
		log:             log,
		storagePostgres: storage.NewStoragePostgres(db),
	}
}
