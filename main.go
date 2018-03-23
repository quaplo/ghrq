package main

import (
	"log"
	"os"
)

// ConfigRabbitMQ connection for rabbitMQ service
type ConfigRabbitMQ struct {
	Host     string `yml:"host"`
	Port     string `yml:"port"`
	Vhost    string `yml:"vhost"`
	User     string `yml:"user"`
	Password string `yml:"password"`
}

// Config global application configuration
type Config struct {
	ConfigRabbitMQ
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	if len(os.Args) == 1 {
		log.Fatal("You must define client type")
	}

	switch os.Args[1] {
	case "c":
		log.Print("Starting AMQP consumer")
		Recive()
	case "p":
		log.Print("Starting AMQP publisher")
		Publish()
	default:
		log.Print("Unknow type")
	}
}
