package application

import (
	"go_boilerplate/config"
	"go_boilerplate/pkg/logger"
	"go_boilerplate/pkg/pubsub"
	"go_boilerplate/storage"

	"github.com/jmoiron/sqlx"
)

// Application ...
type Application struct {
	cfg             config.Config
	log             logger.Logger
	storagePostgres storage.PostgresStorageI
	rmq             *pubsub.RMQ
}

// New ...
func New(cfg config.Config, log logger.Logger, db *sqlx.DB, rmq *pubsub.RMQ) *Application {
	return &Application{
		cfg:             cfg,
		log:             log,
		storagePostgres: storage.NewStoragePostgres(db),
		rmq:             rmq,
	}
}

// RegisterConsumers ...
func (s *Application) RegisterConsumers() {
	s.rmq.AddConsumer(
		"go_boilerplate.application.create", // consumerName
		"application",                       // exchangeName
		"application.create",                // queueName
		"application.create",                // routingKey
		s.createApplicationListener,
	)

	s.rmq.AddConsumer(
		"go_boilerplate.application.created", // consumerName
		"application",                        // exchangeName
		"application.created",                // queueName
		"application.created",                // routingKey
		s.applicationCreatedListener,
	)
}
