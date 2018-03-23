package rabbitMQ

import (
	"log"

	"github.com/streadway/amqp"
)

// Publish simple message publishing
func (rabbit *RabbitMQ) Publish() {

	rabbit.SetQueue()
	defer rabbit.Close()

	for i := 0; i < 1000; i++ {
		body := "hello"
		err := rabbit.Channel.Publish(
			"go_test",         // exchange
			rabbit.Queue.Name, // routing key
			false,             // mandatory
			false,             // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		log.Printf(" [x] Sent %s", body)
		failOnError(err, "Failed to publish a message")
	}
}
