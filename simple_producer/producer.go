package main

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	// connect to RMQ Server
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatalln("Failed to connect to RabbitMQ", err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln("Failed to open a channel", err)
	}

	defer ch.Close()

	//  Declare a queue
	q, err := ch.QueueDeclare(
		"hello", // queue name
		false,   // Durability
		false,   // declare when unused
		false,   // exclusive
		false,   // no wait
		nil,
	)
	if err != nil {
		log.Fatalln("Failed to declare a queue", err)
	}

	err = ch.Publish("", q.Name, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello, RabbitMQ!"),
	})

	if err != nil {
		log.Fatalln("Failed to publish message: ", err)
	}
}
