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

// AMQPQueue queue define
type AMQPQueue struct {
	Name         string
	Durable      bool
	UnusedDelete bool
	Exclusive    bool
	NoWait       bool
	Arguments    amqp.Table
}

// NewRabbitMQ construct
func NewRabbitMQ() (*RabbitMQ, error) {
	// create new instance
	rabbitMQ := RabbitMQ{}
	// define err as error type
	var err error

	// cerate connection
	if rabbitMQ.Conn, err = amqp.Dial(os.Getenv("AMQP_SERVER_LOGIN")); err != nil {
		return nil, err
	}

	// create channel
	if rabbitMQ.Channel, err = rabbitMQ.Conn.Channel(); err != nil {
		return nil, err
	}

	return &rabbitMQ, nil
}

// Close just close connection and channel
func (rabbit *RabbitMQ) Close() {
	rabbit.Conn.Close()
	rabbit.Channel.Close()
}

// SetQueue define current queue
func (rabbit *RabbitMQ) SetQueue(queue AMQPQueue) {
	var err error
	// @toto validate if channel not exist
	if rabbit.Queue, err = rabbit.Channel.QueueDeclare(
		queue.Name,         // name
		queue.Durable,      // durable
		queue.UnusedDelete, // delete when unused
		queue.Exclusive,    // exclusive
		queue.NoWait,       // no-wait
		queue.Arguments,    // arguments
	); err != nil {
		log.Fatal("Canot define queue")
	}
}
