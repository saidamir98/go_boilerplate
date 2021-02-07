package event

import (
	"errors"
	"fmt"

	"github.com/streadway/amqp"
)

// Publisher ...
type Publisher struct {
	channel      *amqp.Channel
	exchangeName string
}

// Push ...
func (p *Publisher) Push(routingKey string, msg amqp.Publishing) error {
	err := p.channel.Publish(
		p.exchangeName,
		routingKey,
		false,
		false,
		msg,
	)
	return err
}

// NewPublisher ...
func (rmq *RMQ) NewPublisher(publisherName, exchangeName string) error {
	if rmq.publishers[publisherName] != nil {
		return errors.New("publisher with the same name already exists: " + publisherName)
	}

	ch, err := rmq.conn.Channel()

	if err != nil {
		return err
	}

	err = declareExchange(ch, exchangeName)

	if err != nil {
		fmt.Printf("Exchange Declare: %s", err.Error())
		return err
	}

	rmq.publishers[publisherName] = &Publisher{
		channel:      ch,
		exchangeName: exchangeName,
	}

	return nil
}

// Close ...
func (p *Publisher) Close() error {
	return p.channel.Close()
}
