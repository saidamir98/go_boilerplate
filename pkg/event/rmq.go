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
		consumerErrors: make(chan error),
		publishers:     make(map[string]*Publisher),
	}

	go rmq.receiveConsumerError()

	return rmq, nil
}

// RunConsumers ...
func (rmq *RMQ) RunConsumers(ctx context.Context) {
	var wg sync.WaitGroup

	for _, consumer := range rmq.consumers {
		c := consumer
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			c.Start(ctx)
		}(&wg)
		fmt.Println("Key:", c.queueName, "=>", "consumer:", c)
	}

	wg.Wait()
}

func (rmq *RMQ) receiveConsumerError() {
	for err := range rmq.consumerErrors {
		rmq.log.Error("consumer error", logger.Error(err))
	}
}
