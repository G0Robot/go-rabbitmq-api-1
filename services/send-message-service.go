package services

import (
	"go-rabbitmq-api-1/models"
	"log"

	"github.com/streadway/amqp"
)

func SendMessageService(message models.Message) {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"demonstration", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	body := message.MessageContent

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
