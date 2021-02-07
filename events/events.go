package events

import (
	"errors"
	"fmt"
	"go_boilerplate/config"
	"go_boilerplate/pkg/event"
	"go_boilerplate/pkg/logger"

	"github.com/jmoiron/sqlx"
	"github.com/streadway/amqp"
)

// New ...
func New(cfg config.Config, log logger.Logger, db *sqlx.DB, amqpConn *amqp.Connection) (*event.RMQ, error) {
	rmq, err := event.NewRMQ(amqpConn, log)
	if err != nil {
		return nil, err
	}

	err = rmq.NewConsumer("#####.create", "#####", "#####.create", "#####.create", func(delivery amqp.Delivery) error {
		fmt.Println("celebrity.create")
		fmt.Println(delivery)

		return errors.New("some error")
	})
	if err != nil {
		return nil, err
	}

	err = rmq.NewConsumer("#####.update", "#####", "#####.update", "#####.update", func(delivery amqp.Delivery) error {
		fmt.Println("celebrity.update")
		fmt.Println(delivery)
		return nil
	})
	if err != nil {
		return nil, err
	}

	// err = rmq.NewConsumer("#####.update", "#####", "#####.update", "#####.update", func(delivery amqp.Delivery) error {
	// 	fmt.Println("celebrity.update")
	// 	fmt.Println(delivery)
	// 	return nil
	// })
	// if err != nil {
	// 	return nil, err
	// }

	err = rmq.NewConsumer("#####.delete", "#####", "#####.delete", "#####.delete", func(delivery amqp.Delivery) error {
		fmt.Println("celebrity.delete")
		fmt.Println(delivery)
		return nil
	})
	if err != nil {
		return nil, err
	}

	fmt.Println(rmq)

	return rmq, err
}
