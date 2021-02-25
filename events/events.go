package events

import (
	"encoding/json"
	"fmt"
	"go_boilerplate/config"
	"go_boilerplate/events/application"
	"go_boilerplate/go_boilerplate_modules/application_service"
	"go_boilerplate/pkg/event"
	"go_boilerplate/pkg/logger"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/streadway/amqp"
)

// New ...
func New(cfg config.Config, log logger.Logger, db *sqlx.DB, amqpURI string) (*event.RMQ, error) {
	rmq, err := event.NewRMQ(amqpURI, log)
	if err != nil {
		return nil, err
	}

	go func() {
		time.Sleep(time.Second * 7)
		conn, err := amqp.Dial(amqpURI)
		if err != nil {
			fmt.Println(err)
		}
		publisher := event.NewPublisher(conn, "exchange.application.v1")

		for i := 0; i < 1000; i++ {
			uuid, _ := uuid.NewRandom()
			entity := application_service.CreateApplicationModel{
				ID:   uuid.String(),
				Body: fmt.Sprint(i),
			}

			b, err := json.Marshal(entity)

			err = publisher.Push("application.create", amqp.Publishing{
				ReplyTo:       "application.update",
				CorrelationId: fmt.Sprint(i),
				Body:          b,
			})
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(time.Second * 1)
		}

	}()

	applicationService := application.New(cfg, log, db)

	rmq.AddConsumer(
		"go_boilerplate", // ConsumerName: {consuming service}
		"application",     // ExchangeName: model => exchange sending only events or commands about application
		"application.event.created",    // QueueName: {model}.{type:(event|command)}.{event-name(model.actioned)|command(event-name)}
		"application.event.created",          // routingKey:  {model}.{type:(event|command)}.{event-name(model.actioned)|command(event-name)}
		applicationService.CreateApplicationListener,
	)

	rmq.AddConsumer(
		"go_boilerplate", // ConsumerName: {consuming service}
		"application",     // ExchangeName: model => exchange sending only events or commands about application
		"application.event.updated",    // QueueName: {model}.{type:(event|command)}.{event-name(model.actioned)|command(event-name)}
		"application.event.updated",    // routingKey:  {model}.{type:(event|command)}.{event-name(model.actioned)|command(event-name)}
		applicationService.UpdateApplicationListener,
	)

	fmt.Println(rmq)

	return rmq, err
}
