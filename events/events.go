package events

import (
	"context"
	"go_boilerplate/config"
	"go_boilerplate/events/application"
	"go_boilerplate/pkg/logger"
	"go_boilerplate/pkg/pubsub"

	"github.com/jmoiron/sqlx"
)

// PubsubServer ...
type PubsubServer struct {
	cfg config.Config
	log logger.Logger
	db  *sqlx.DB
	RMQ *pubsub.RMQ
}

// New ...
func New(cfg config.Config, log logger.Logger, db *sqlx.DB) (*PubsubServer, error) {
	rmq, err := pubsub.NewRMQ(cfg.RabbitURI, log)
	if err != nil {
		return nil, err
	}

	// Register publishers here -------->
	rmq.AddPublisher("application") // one publisher is enough for application service
	// <--------

	// // test publisher ------------>
	// go func() {
	// 	time.Sleep(time.Millisecond * 3000)
	// 	for i := 0; i < 10000; i++ {
	// 		uuid, _ := uuid.NewRandom()
	// 		entity := application_service.CreateApplicationModel{
	// 			ID:   uuid.String(),
	// 			Body: fmt.Sprint(i),
	// 		}

	// 		b, err := json.Marshal(entity)

	// 		fmt.Println("---------------------------------------------------------------------")
	// 		err = rmq.Push("application", "application.create", amqp.Publishing{
	// 			ContentType:   "application/json",
	// 			DeliveryMode:  amqp.Persistent,
	// 			ReplyTo:       "application.created",
	// 			CorrelationId: fmt.Sprint(i),
	// 			Body:          b,
	// 		})

	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}

	//		// time.Sleep(time.Millisecond * 1000)
	// 	}
	// }()
	// // <----------

<<<<<<< HEAD
	return &PubsubServer{
		cfg: cfg,
		log: log,
		db:  db,
		RMQ: rmq,
	}, nil
}
=======
	rmq.AddConsumer(
		"go_boilerplate", // ConsumerName: {consuming service}
		"application",     // ExchangeName: model => exchange sending only events or commands about application
		"application.event.created",    // QueueName: {model}.{type:(event|command)}.{event-name(model.actioned)|command(event-name)}
		"application.event.created",          // routingKey:  {model}.{type:(event|command)}.{event-name(model.actioned)|command(event-name)}
		applicationService.CreateApplicationListener,
	)

	rmq.AddConsumer(
		"go_boilerplate", // ConsumerName: {consuming service}
		"application",     // ExchangeName: model => exchange sending only events or commands about application
		"application.event.updated",    // QueueName: {model}.{type:(event|command)}.{event-name(model.actioned)|command(event-name)}
		"application.event.updated",    // routingKey:  {model}.{type:(event|command)}.{event-name(model.actioned)|command(event-name)}
		applicationService.UpdateApplicationListener,
	)
>>>>>>> cc84aa670aea59a8827d78466b26456a9da6d02d

// Run ...
func (s *PubsubServer) Run(ctx context.Context) {
	applicationServer := application.New(s.cfg, s.log, s.db, s.RMQ)
	applicationServer.RegisterConsumers()

	s.RMQ.RunConsumers(ctx)
}
