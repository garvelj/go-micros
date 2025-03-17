package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name of the exchange
		"topic",      // type of the topic
		true,         // is it durable
		false,        // do you get rid of it when you're done with it
		false,        // internal? no between our microservices
		false,        // no-wait?
		nil,          // additional args?
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name (pick your own name)
		false, // durable?
		false, // delete when unused?
		true,  // exclusive
		false, //nowait?
		nil,   //args?
	)
}
