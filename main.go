package main

import (
	"log"
	"os"

	"github.com/quaplo/ghrq/rabbitMQ"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatal("You must define client type")
	}

	rabbit, err := rabbitMQ.NewRabbitMQ()
	if err != nil {
		log.Fatalln(err)
	}

	switch os.Args[1] {
	case "c":
		log.Print("Starting AMQP consumer")
		rabbit.Recive()
	case "p":
		log.Print("Starting AMQP publisher")
		rabbit.Publish()
	default:
		log.Print("Unknow type")
	}
}
