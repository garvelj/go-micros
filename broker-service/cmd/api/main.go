package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const port = "80"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {
	rabbitConn, err := connect()
	if err != nil {
		os.Exit(1)
	}

	defer rabbitConn.Close()

	app := Config{
		Rabbit: rabbitConn,
	}

	log.Printf("Starting broker service on port %s", port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

// COPIED FROM LISTENER SERVICE
func connect() (*amqp.Connection, error) {
	var (
		counts     int
		backOff    = 1 * time.Second
		connection *amqp.Connection
	)

	for {
		// Attempt connection
		// this has to match what is in the docker-compose file
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
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
