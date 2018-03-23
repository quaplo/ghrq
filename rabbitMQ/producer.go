package rabbitMQ

import (
	"log"

	"github.com/streadway/amqp"
)

func Publish() {
	conn, err := amqp.Dial("amqp://wtzlveww:lrbUQufW88S66j6fifVic9Pv6nXHLGuB@skunk.rmq.cloudamqp.com:5672/wtzlveww")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"go_test", // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	failOnError(err, "Failed to declare a queue")

	for i := 0; i < 10; i++ {
		body := "hello"
		err = ch.Publish(
			"go_test", // exchange
			q.Name,    // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		log.Printf(" [x] Sent %s", body)
		failOnError(err, "Failed to publish a message")
	}

}
