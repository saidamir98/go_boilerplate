package pubsub

import "github.com/streadway/amqp"

func declareQueue(ch *amqp.Channel, name string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

// func declareExchange(ch *amqp.Channel, name string) error {
// 	return ch.ExchangeDeclare(
// 		name,               // name
// 		amqp.ExchangeTopic, // type
// 		true,               // durable
// 		false,              // auto-deleted
// 		false,              // internal
// 		false,              // no-wait
// 		nil,                // arguments
// 	)
// }
