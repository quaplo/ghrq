package rabbitMQ

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

// RabbitMQ global RabbitMQ acces
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}

// NewRabbitMQ construct
func NewRabbitMQ() (*RabbitMQ, error) {
	// create new instance
	rabbitMQ := RabbitMQ{}

	// cerate connection
	conn, err := amqp.Dial(os.Getenv("AMQP_SERVER_LOGIN"))
	if err != nil {
		return nil, err
	}
	rabbitMQ.Conn = conn

	// create channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	rabbitMQ.Channel = ch

	return &rabbitMQ, nil
}

// Close just close connection and channel
func (rabbit *RabbitMQ) Close() {
	rabbit.Conn.Close()
	rabbit.Channel.Close()
}

// SetQueue define current queue
func (rabbit *RabbitMQ) SetQueue() {
	// @toto validate if channel not exist
	queue, err := rabbit.Channel.QueueDeclare(
		"go_test", // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		log.Fatal("Canot define queue")
	}

	rabbit.Queue = queue
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
