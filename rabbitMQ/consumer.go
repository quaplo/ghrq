package rabbitMQ

import (
	"log"
)

// Recive recive messages
func (rabbit *RabbitMQ) Recive() {

	rabbit.SetQueue()
	defer rabbit.Close()

	msgs, err := rabbit.Channel.Consume(
		rabbit.Queue.Name, // queue
		"i_am_consumer",   // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)

	if err != nil {
		log.Fatalln(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
