package rabbitMQ

import (
	"github.com/streadway/amqp"
)

// RabbitMQ global RabbitMQ acces
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}
