package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wawayes/bi-chatgpt-golang/pkg/logx"
)

func ReceiveFromRabbitMQ() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672")
	failOnError(err, "Failed to connect to rabbitmq")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	msg, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msg {
			logx.Info("Received a message: " + string(d.Body))
		}
	}()

	logx.Info(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
