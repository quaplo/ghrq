package rabbitMQ

import (
	"log"

	"github.com/streadway/amqp"
)

// AMQPMessage format for sending message
type AMQPMessage struct {
	Exchange    string
	RoutingKey  string
	Mandatory   bool
	Immediate   bool
	ContentType string
	Body        string
}

// Publish simple message publishing
func (rabbit *RabbitMQ) Publish(message AMQPMessage) {
	defer rabbit.Close()

	if err := rabbit.Channel.Publish(
		message.Exchange,   // exchange
		message.RoutingKey, // routing key
		message.Mandatory,  // mandatory
		message.Immediate,  // immediate
		amqp.Publishing{
			ContentType: message.ContentType,  // content type
			Body:        []byte(message.Body), // body
		}); err != nil {
		log.Fatal("Failed to publish a message")
	} else {
		log.Printf(" [x] Sent %s", message.Body)
	}
}
