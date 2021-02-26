package pubsub

import (
	"errors"
	"go_boilerplate/pkg/logger"

	"github.com/streadway/amqp"
)

// Publisher ...
type Publisher struct {
	channel      *amqp.Channel
	exchangeName string
}

// AddPublisher ...
func (rmq *RMQ) AddPublisher(exchangeName string) {
	if rmq.publishers[exchangeName] != nil {
		rmq.log.Warn("publisher exists", logger.Error(errors.New("publisher with the same exchange name already exists: "+exchangeName)))
		return
	}

	ch, err := rmq.conn.Channel()

	if err != nil {
		panic(err)
	}

	// err = declareExchange(ch, exchangeName)

	// if err != nil {
	// 	fmt.Printf("Exchange Declare: %s", err.Error())
	// 	panic(err)
	// }

	rmq.publishers[exchangeName] = &Publisher{
		channel:      ch,
		exchangeName: exchangeName,
	}
}

// Push ...
func (rmq *RMQ) Push(exchangeName string, routingKey string, msg amqp.Publishing) error {
	p := rmq.publishers[exchangeName]

	if p == nil {
		return errors.New("publisher with that exchange name doesn't exists: " + exchangeName)
	}

	err := p.channel.Publish(
		p.exchangeName,
		routingKey,
		false,
		false,
		msg,
	)
	return err
}

// // PushRPC ...
// func (rmq *RMQ) PushRPC(exchangeName string, routingKey string, msg amqp.Publishing) error {

// 	return nil
// }

// Close ...
func (p *Publisher) Close() error {
	return p.channel.Close()
}
