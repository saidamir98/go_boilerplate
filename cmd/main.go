package main

import (
	"context"
	"fmt"
	"go_boilerplate/api"
	"go_boilerplate/config"
	"go_boilerplate/events"
	"go_boilerplate/pkg/logger"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// var wg sync.WaitGroup

	// wg.Add(1)
	// go func(wg *sync.WaitGroup) {
	// 	defer wg.Done()
	// 	eventServer, err := events.New(cfg, log, db, amqpConn)
	// 	if err != nil {
	// 		log.Panic("error on the event server", logger.Error(err))
	// 		return
	// 	}
	// 	eventServer.RunConsumers(context.Background()) // should be some recovery
	// 	log.Panic("event server has finished")
	// }(&wg)

	// wg.Add(1)
	// go func(wg *sync.WaitGroup) {
	// 	defer wg.Done()
	// 	apiServer := api.New(cfg, log, db)
	// 	err := apiServer.Run(cfg.HTTPPort)
	// 	if err != nil {
	// 		log.Panic("error on the api server", logger.Error(err))
	// 		return
	// 	}
	// 	log.Panic("api server has finished")
	// }(&wg)

	// wg.Wait()

	pubsubServer, err := events.New(cfg, log, db)
	if err != nil {
		log.Panic("error on the event server", logger.Error(err))
	}

	apiServer, err := api.New(cfg, log, db)
	if err != nil {
		log.Panic("error on the api server", logger.Error(err))
	}

	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		pubsubServer.Run(ctx) // it should run forever if there is any consumer
		log.Panic("event server has finished")
		return nil
	})

	group.Go(func() error {
		err := apiServer.Run(cfg.HTTPPort) // this method will block the calling goroutine indefinitely unless an error happens
		if err != nil {
			panic(err)
		}
		log.Panic("api server has finished")
		return nil
	})

	err = group.Wait()
	if err != nil {
		log.Panic("error on the server", logger.Error(err))
	}
}
