package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name of the exchange
		"topic",      // what topic are we going to use
		true,         // is the exchange durable
		false,        // auto delete?
		false,        // this exchange isn't used just internally
		false,        // no wait, don't worry about this one too much now
		nil,          // no specific arguments
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"randomName", // name of the queue
		false,        // durable?
		false,        // delete when unused?
		true,         // exclusive
		false,        // no-wait?
		nil,          // arguments
	)
}
