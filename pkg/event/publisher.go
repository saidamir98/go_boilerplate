package event

import (
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
func NewPublisher(conn *amqp.Connection, exchangeName string) Publisher {
	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	// err = declareExchange(ch, exchangeName)

	// if err != nil {
	// 	fmt.Printf("Exchange Declare: %s", err.Error())
	// 	panic(err)
	// }

	return Publisher{
		channel:      ch,
		exchangeName: exchangeName,
	}
}

// Close ...
func (p *Publisher) Close() error {
	return p.channel.Close()
}
