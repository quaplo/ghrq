package main

import (
	"log"
	"os"

	"github.com/quaplo/ghrq/rabbitMQ"
)

func main() {

	// validation args
	if len(os.Args) == 1 {
		log.Fatal("You must define client type")
	}

	// define rabbitMQ service
	rabbit, err := rabbitMQ.NewRabbitMQ()
	if err != nil {
		log.Fatalln(err)
	}

	// define queue
	queue := rabbitMQ.AMQPQueue{
		Name:         "go_test",
		Durable:      true,
		UnusedDelete: false,
		Exclusive:    false,
		NoWait:       false,
	}
	// set queue default for this example
	rabbit.SetQueue(queue)

	switch os.Args[1] {
	case "c":
		log.Print("Starting AMQP consumer")
		rabbit.Recive()
	case "p":
		log.Print("Starting AMQP publisher")

		// define new message
		message := rabbitMQ.AMQPMessage{
			Exchange:    "go_test",
			RoutingKey:  "go_test",
			Mandatory:   false,
			Immediate:   false,
			ContentType: "text/plain",
			Body:        "say hello again",
		}

		rabbit.Publish(message)
	default:
		log.Print("Unknow type")
	}
}
