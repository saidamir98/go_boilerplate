package event

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

// Consumer ...
type Consumer struct {
	consumerName string
	channel      *amqp.Channel
	exchangeName string
	routingKey   string
	queueName    string
	handler      func(amqp.Delivery) error
	messages     <-chan amqp.Delivery
	errors       chan error
}

// NewConsumer ...
func (rmq *RMQ) NewConsumer(consumerName, exchangeName, routingKey, queueName string, handler func(amqp.Delivery) error) error {
	if rmq.consumers[consumerName] != nil {
		return errors.New("consumer with the same name already exists: " + consumerName)
	}

	ch, err := rmq.conn.Channel()

	if err != nil {
		return err
	}

	// err = declareExchange(ch, exchangeName)

	// if err != nil {
	// 	fmt.Printf("Exchange Declare: %s", err.Error())
	// 	return err
	// }

	q, err := declareQueue(ch, queueName)

	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,
		routingKey,
		exchangeName,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	// messages, err := ch.Consume(
	// 	queueName,
	// 	consumerName,
	// 	false,
	// 	false,
	// 	false,
	// 	true,
	// 	nil,
	// )

	// if err != nil {
	// 	return err
	// }

	rmq.consumers[consumerName] = &Consumer{
		consumerName: consumerName,
		channel:      ch,
		exchangeName: exchangeName,
		routingKey:   routingKey,
		queueName:    queueName,
		handler:      handler,
		errors:       rmq.consumerErrors,
	}

	return nil
}

// Start ...
func (c *Consumer) Start(ctx context.Context) {
	var err error
	c.messages, err = c.channel.Consume(
		c.queueName,
		c.consumerName,
		false,
		false,
		false,
		true,
		nil)
	if err != nil {
		c.errors <- err
		return
	}
	for {
		select {
		case msg, ok := <-c.messages:
			if !ok {
				fmt.Println(c.queueName)
				c.errors <- errors.New("error while reading consumer messages")
				time.Sleep(time.Duration(5000 * time.Millisecond))
			} else {
				fmt.Println("rmq -> msg ", msg)
				err := c.handler(msg)
				if err != nil {
					c.errors <- err
				}
			}
		case <-ctx.Done():
			{
				err := c.channel.Cancel("", true)
				if err != nil {
					c.errors <- err
				}
				return
			}
		}
	}
}
