package storage

import (
	"go_boilerplate/storage/postgres"
	"go_boilerplate/storage/repo"

	"github.com/jmoiron/sqlx"
)

// PostgresStorageI ...
type PostgresStorageI interface {
	Application() repo.ApplicationStorageI
}

type storagePostgres struct {
	db              *sqlx.DB
	applicationRepo repo.ApplicationStorageI
}

// NewStoragePostgres ...
func NewStoragePostgres(db *sqlx.DB) PostgresStorageI {
	return &storagePostgres{
		db:              db,
		applicationRepo: postgres.NewApplicationRepo(db),
	}
}

// Application ...
func (s storagePostgres) Application() repo.ApplicationStorageI {
	return s.applicationRepo
}


