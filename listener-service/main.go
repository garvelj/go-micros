package main

// Listener service is a service that talks to Rabbit MQ
// which is a server that manages queues

// What is the actual use case:
// Someone would want to authenticate,
// And the broker won't communicate with authentication service directly,
// Instead, it pushes instructions into RabbitMQ
// Which is a AMQP - Advanced Messaging Queue Protocol Server
// RabbitMQ would take that request and add it to the queue
// And the listener pulls from that queue
// And and calls the apropriate service based on that.

// So what happens in real life:
// Broker recieves a request
// Says "cool, ill just push that to the queue"
// Then he forgets about it
// The listener service gets that request from the queue and says
// "what should i do with this? should i log something, send an email..."
// In this case it will be the authentication of a user.

// This sounds complicated but actually isn't.

// The driver that is needed for this communication with RabbitMQ
// github.com/rabbitmq/amqp091-go

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// First we'll try to connect with RabbitMQ
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()

	// Then start listening for messages
	// This service won't be fetching the messages
	// Instead, the queue will push it right to us
	// We'll listen for certain queues, and once there's an event on those
	// We get it from the queue

	// Create a consumer

	// Watch the queue and consume the events
}

// So we obviously we have to add the listener service to our docker compose file
// We have to see an available rabbitmq docker image on docker hub
// -rc in the tag of the image stands for release candidate, and i'ts usually not safe to use them

// Let's put the connection to the RabbitMQ in it's own function
func connect() (*amqp.Connection, error) {
	var (
		counts     int
		backOff    = 1 * time.Second
		connection *amqp.Connection
	)

	for {
		// Attempt connection
		c, err := amqp.Dial("amqp://guest:guest@localhost")
		if err != nil {

			// if error occured -> print and increase count
			fmt.Println("RabbitMQ not ready yet...")
			fmt.Println("Err: ", err)
			counts++
		} else {

			// if connection is successful, break from the loop
			connection = c
			break
		}

		// we don't want this to be looping forever
		// So after a while if it isn't connected
		// There's clearly a problem
		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		fmt.Println("Backing off...")
		time.Sleep(backOff)
	}

	return connection, nil
}
