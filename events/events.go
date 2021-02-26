package events

import (
	"context"
	"encoding/json"
	"fmt"
	"go_boilerplate/config"
	"go_boilerplate/events/application"
	"go_boilerplate/go_boilerplate_modules/application_service"
	"go_boilerplate/pkg/logger"
	"go_boilerplate/pkg/pubsub"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/streadway/amqp"
)

// PubsubServer ...
type PubsubServer struct {
	cfg config.Config
	log logger.Logger
	db  *sqlx.DB
	rmq *pubsub.RMQ
}

// New ...
func New(cfg config.Config, log logger.Logger, db *sqlx.DB) (*PubsubServer, error) {
	rmq, err := pubsub.NewRMQ(cfg.RabbitURI, log)
	if err != nil {
		return nil, err
	}

	rmq.AddPublisher("application")

	// test publisher ------------>
	go func() {
		time.Sleep(time.Millisecond * 3000)
		for i := 0; i < 1000; i++ {
			uuid, _ := uuid.NewRandom()
			entity := application_service.CreateApplicationModel{
				ID:   uuid.String(),
				Body: fmt.Sprint(i),
			}

			b, err := json.Marshal(entity)

			fmt.Println("---------------------------------------------------------------------")
			err = rmq.Push("application", "application.create", amqp.Publishing{
				ContentType:   "application/json",
				DeliveryMode:  amqp.Persistent,
				ReplyTo:       "application.created",
				CorrelationId: fmt.Sprint(i),
				Body:          b,
			})

			if err != nil {
				fmt.Println(err)
			}

			time.Sleep(time.Millisecond * 1000)
		}
	}()
	// <----------

	return &PubsubServer{
		cfg: cfg,
		log: log,
		db:  db,
		rmq: rmq,
	}, nil
}

// Run ...
func (s *PubsubServer) Run(ctx context.Context) {
	applicationServer := application.New(s.cfg, s.log, s.db, s.rmq)
	applicationServer.RegisterConsumers()

	s.rmq.RunConsumers(ctx)
}
