package mq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wawayes/bi-chatgpt-golang/pkg/logx"
	"time"
)

func SendToRabbitMQ() {
	// 连接rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 开放channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明队列
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// cancelFunc 告诉操作放弃该工作，不会等待工作停止
	defer cancel()

	body := "hello world"

	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(body)})
	failOnError(err, "Failed to publish a message")

	logx.Info("[x] Sent : " + body)

}

func failOnError(err error, msg string) {
	if err != nil {
		logx.Error(err.Error() + ": " + msg)
	}
}
