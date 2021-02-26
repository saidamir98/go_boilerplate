package application

import (
	"encoding/json"
	"fmt"
	"go_boilerplate/go_boilerplate_modules/application_service"
	"go_boilerplate/pkg/logger"

	"github.com/streadway/amqp"
)

// createApplicationListener ...
func (s *Application) createApplicationListener(delivery amqp.Delivery) error {
	var (
		entity application_service.CreateApplicationModel
	)

	err := json.Unmarshal(delivery.Body, &entity)
	if err != nil {
		s.log.Error("unmarshal error", logger.Any("[]byte", delivery.Body), logger.Any("error", err))
		delivery.Nack(false, false)
		return err
	}

	fmt.Println(entity)

	res, err := s.storagePostgres.Application().Create(entity)
	if err != nil {
		s.log.Error("storage error", logger.Any("entity", entity), logger.Any("error", err))
		delivery.Nack(false, false)
		return err
	}

	b, err := json.Marshal(res)
	if err != nil {
		s.log.Error("marshal error", logger.Any("struct", res), logger.Any("error", err))
		delivery.Nack(false, false)
		return err
	}

	resp := amqp.Publishing{
		ContentType:   "application/json",
		DeliveryMode:  amqp.Persistent,
		CorrelationId: delivery.CorrelationId,
		Body:          b,
	}

	err = s.rmq.Push("application", delivery.ReplyTo, resp)
	if err != nil {
		s.log.Error("publish error", logger.Any("resp", resp), logger.Any("error", err))
		delivery.Nack(false, false)
		return err
	}

	delivery.Ack(false)
	return nil
}

// applicationCreatedListener ...
func (s *Application) applicationCreatedListener(delivery amqp.Delivery) error {
	var (
		entity application_service.ApplicationCreatedModel
	)

	err := json.Unmarshal(delivery.Body, &entity)
	if err != nil {
		s.log.Error("unmarshal error", logger.Any("[]byte", delivery.Body), logger.Any("error", err))
		delivery.Nack(false, false)
		return err
	}

	fmt.Println(entity)

	delivery.Ack(false)
	return nil
}
