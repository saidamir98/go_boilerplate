package application

import (
	"encoding/json"
	"fmt"
	"go_boilerplate/go_boilerplate_modules/application_service"
	"go_boilerplate/pkg/logger"

	"github.com/streadway/amqp"
)

// CreateApplicationListener ...
func (s *Application) CreateApplicationListener(delivery amqp.Delivery) (resp amqp.Publishing, err error) {
	var (
		entity application_service.CreateApplicationModel
	)

	resp.CorrelationId = delivery.CorrelationId

	err = json.Unmarshal(delivery.Body, &entity)
	if err != nil {
		s.log.Error("unmarshal error", logger.Any("[]byte", delivery.Body), logger.Any("error", err))
		delivery.Nack(false, false)
		return
	}

	fmt.Println(entity)

	res, err := s.storagePostgres.Application().Create(entity)
	if err != nil {
		s.log.Error("storage error", logger.Any("entity", entity), logger.Any("error", err))
		delivery.Nack(false, false)
		return
	}

	resp.Body, err = json.Marshal(res)
	if err != nil {
		s.log.Error("marshal error", logger.Any("struct", res), logger.Any("error", err))
		delivery.Nack(false, false)
		return
	}

	delivery.Ack(false)
	return
}

// UpdateApplicationListener ...
func (s *Application) UpdateApplicationListener(delivery amqp.Delivery) (resp amqp.Publishing, err error) {
	var (
		entity application_service.CreateApplicationModel
	)

	resp = amqp.Publishing{
		CorrelationId: delivery.CorrelationId,
	}

	err = json.Unmarshal(delivery.Body, &entity)
	if err != nil {
		s.log.Error("unmarshal error", logger.Any("[]byte", delivery.Body), logger.Any("error", err))
		delivery.Nack(false, false)
	}

	fmt.Println("---------->")
	fmt.Println(entity)
	fmt.Println("----------<")

	// res, err := s.storagePostgres.Application().Create(entity)

	// resp.Body, err = json.Marshal(res)
	// if err != nil {
	// 	s.log.Error("marshal error", logger.Any("struct", res), logger.Any("error", err))
	// 	delivery.Nack(false, false)
	// 	return
	// }

	delivery.Nack(false, false)
	return
}
