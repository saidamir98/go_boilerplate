package event

import (
	"context"
	"fmt"
	"go_boilerplate/pkg/logger"
	"sync"

	"github.com/streadway/amqp"
)

// RMQ ...
type RMQ struct {
	conn           *amqp.Connection
	log            logger.Logger
	consumers      map[string]*Consumer
	consumerErrors chan error
	publishers     map[string]*Publisher
}

// NewRMQ ...
func NewRMQ(conn *amqp.Connection, log logger.Logger) (*RMQ, error) {
	rmq := &RMQ{
		conn:           conn,
		log:            log,
		consumers:      make(map[string]*Consumer),
		consumerErrors: make(chan error), // must be buffered size
		publishers:     make(map[string]*Publisher),
	}

	go rmq.receiveConsumerError()

	return rmq, nil
}

// RunConsumers ...
func (rmq *RMQ) RunConsumers(ctx context.Context) {
	var wg sync.WaitGroup

	for _, consumer := range rmq.consumers {
		wg.Add(1)
		go func(wg *sync.WaitGroup, c *Consumer) {
			defer wg.Done()
			c.Start(ctx)
		}(&wg, consumer)
		fmt.Println("Key:", consumer.queueName, "=>", "consumer:", consumer)
	}

	wg.Wait()
}

func (rmq *RMQ) receiveConsumerError() {
	for err := range rmq.consumerErrors {
		rmq.log.Error("consumer error", logger.Error(err))
	}
}
