package pubsub

import (
	"context"
	"errors"

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

// AddConsumer ...
func (rmq *RMQ) AddConsumer(consumerName, exchangeName, queueName, routingKey string, handler func(amqp.Delivery) error) {
	if rmq.consumers[consumerName] != nil {
		panic(errors.New("consumer with the same name already exists: " + consumerName))
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

	q, err := declareQueue(ch, queueName)

	if err != nil {
		panic(err)
	}

	err = ch.QueueBind(
		q.Name,
		routingKey,
		exchangeName,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	messages, err := ch.Consume(
		queueName,
		consumerName,
		false,
		false,
		false,
		true,
		nil,
	)

	if err != nil {
		panic(err)
	}

	rmq.consumers[consumerName] = &Consumer{
		consumerName: consumerName,
		channel:      ch,
		exchangeName: exchangeName,
		routingKey:   routingKey,
		queueName:    queueName,
		handler:      handler,
		messages:     messages,
		errors:       rmq.consumerErrors,
	}

	return
}

// Start ...
func (c *Consumer) Start(ctx context.Context) {
	// var err error
	// c.messages, err = c.channel.Consume(
	// 	c.queueName,
	// 	c.consumerName,
	// 	false,
	// 	false,
	// 	false,
	// 	true,
	// 	nil)
	// if err != nil {
	// 	c.errors <- err
	// 	return
	// }
	for {
		select {
		case msg, ok := <-c.messages:
			if !ok {
				panic(errors.New("error while reading consumer messages"))
			} else {
				err := c.handler(msg)
				if err != nil {
					c.errors <- err
				}
				// else {
				// 	// if msg.CorrelationId != "" {
				// 	c.pushReplay(msg.ReplyTo, resp)
				// 	// }
				// }
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

// // PushReplay ...
// func (c *Consumer) pushReplay(replyTo string, msg amqp.Publishing) {

// 	err := c.channel.Publish(
// 		c.exchangeName,
// 		replyTo,
// 		false,
// 		false,
// 		msg,
// 	)

// 	if err != nil {
// 		c.errors <- err
// 	}

// 	return
// }
