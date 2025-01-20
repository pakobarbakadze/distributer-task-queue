package queue

import (
	"log"

	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
var Channel *amqp.Channel
var Queue amqp.Queue

func InitQueue() {
	var err error
	Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	Channel, err = Conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	Queue, err = Channel.QueueDeclare(
		"task_queue",
		true,  // Durable
		false, // Auto-delete
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	log.Println("RabbitMQ initialized successfully!")
}

func PublishTask(payload string) error {
	return Channel.Publish(
		"",
		Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		},
	)
}
