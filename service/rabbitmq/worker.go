package main

import (
	"bytes"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wawayes/bi-chatgpt-golang/service/mq"
	"log"
	"time"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672")
	mq.FailOnError(err, "Failed to connect to rabbitmq")
	defer conn.Close()

	ch, err := conn.Channel()
	mq.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	mq.FailOnError(err, "Failed to declare a queue")

	msg, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	mq.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msg {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
		}
	}()

	log.Printf(" [*] Waiting for messages, To exit press CTRL + C")
	<-forever

}
