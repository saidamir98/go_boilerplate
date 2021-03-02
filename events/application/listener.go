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

	res, err := s.storagePostgres.Application().Create(entity)
	if err != nil {
		s.log.Error("storage error", logger.Any("entity", entity), logger.Any("error", err))
		delivery.Nack(false, false)
		return err
	}

	s.log.Info("application has been created", logger.Any("entity", entity), logger.Any("res", res))

	// if it replays result back ------->
	b, err := json.Marshal(res)
	if err != nil {
		s.log.Error("marshal error", logger.Any("struct", res), logger.Any("error", err))
		delivery.Nack(false, false)
		return err
	}

	if len(delivery.ReplyTo) > 0 && len(delivery.CorrelationId) > 0 {
		resp := amqp.Publishing{
			ContentType:   "application/json",
			DeliveryMode:  amqp.Persistent,
			CorrelationId: delivery.CorrelationId,
			Body:          b,
		}

		err = s.rmq.Push("application", delivery.ReplyTo, resp)
		if err != nil {
			s.log.Error(
				"publish error",
				logger.String("exchange", "application"),
				logger.String("routing", delivery.ReplyTo),
				logger.Any("msg", resp),
				logger.Any("body", res),
				logger.Any("error", err),
			)
			delivery.Nack(false, false)
			return err
		}
	}
	// <--------

	delivery.Ack(false)
	return nil
}

// applicationCreatedListener consumes replies from application.create, currently it does nothing, only reads body and prints it
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

// updateApplicationListener ...
func (s *Application) updateApplicationListener(delivery amqp.Delivery) error {
	var (
		entity application_service.UpdateApplicationModel
	)

	err := json.Unmarshal(delivery.Body, &entity)
	if err != nil {
		s.log.Error("unmarshal error", logger.Any("[]byte", delivery.Body), logger.Any("error", err))
		delivery.Nack(false, false)
		return err
	}

	res, err := s.storagePostgres.Application().Update(entity)
	if err != nil {
		s.log.Error("storage error", logger.Any("entity", entity), logger.Any("error", err))
		delivery.Nack(false, false)
		return err
	}

	s.log.Info("application has been updated", logger.Any("entity", entity), logger.Any("res", res))

	delivery.Ack(false)
	return nil
}

// deleteApplicationListener ...
func (s *Application) deleteApplicationListener(delivery amqp.Delivery) error {
	var (
		entity application_service.DeleteApplicationModel
	)

	err := json.Unmarshal(delivery.Body, &entity)
	if err != nil {
		s.log.Error("unmarshal error", logger.Any("[]byte", delivery.Body), logger.Any("error", err))
		delivery.Nack(false, false)
		return err
	}

	res, err := s.storagePostgres.Application().Delete(entity)
	if err != nil {
		s.log.Error("storage error", logger.Any("entity", entity), logger.Any("error", err))
		delivery.Nack(false, false)
		return err
	}

	s.log.Info("application has been updated", logger.Any("entity", entity), logger.Any("res", res))

	delivery.Ack(false)
	return nil
}
