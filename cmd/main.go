package main

import (
	"context"
	"fmt"
	"go_boilerplate/api"
	"go_boilerplate/config"
	"go_boilerplate/events"
	"go_boilerplate/pkg/logger"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.App, cfg.LogLevel)

	psqlConnString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	db, err := sqlx.Connect("postgres", psqlConnString)

	if err != nil {
		log.Panic("error connecting to postgres", logger.Error(err))
	}

	amqpConn, err := amqp.Dial(cfg.RabbitURL)

	if err != nil {
		log.Panic("error connecting to rabbit", logger.Error(err))
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		eventServer, err := events.New(cfg, log, db, amqpConn)
		if err != nil {
			log.Panic("error on the event server", logger.Error(err))
			return
		}
		eventServer.RunConsumers(context.Background())
		log.Panic("event server has finished")
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		apiServer := api.New(cfg, log, db)
		err := apiServer.Run(cfg.HTTPPort)
		if err != nil {
			log.Panic("error on the api server", logger.Error(err))
			return
		}
		log.Panic("api server has finished")
	}(&wg)

	wg.Wait()

	// group, ctx := errgroup.WithContext(context.Background())

	// group.Go(func() error {
	// 	eventServer, err := events.New(cfg, log, db, amqpConn)
	// 	if err != nil {
	// 		return err
	// 		// panic(err)
	// 	}
	// 	eventServer.RunConsumers(ctx)
	// 	log.Panic("event server has finished")
	// 	return nil
	// })

	// group.Go(func() error {
	// 	apiServer := api.New(cfg, log, db)
	// 	err := apiServer.Run(cfg.HTTPPort)
	// 	if err != nil {
	// 		return err
	// 		// panic(err)
	// 	}
	// 	log.Panic("api server has finished")
	// 	return nil
	// })

	// err = group.Wait()
	// if err != nil {
	// 	log.Panic("error on the server", logger.Error(err))
	// }
}
