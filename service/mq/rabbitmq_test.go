package mq

import "testing"

func Test_RabbitMQ(test *testing.T) {
	SendToRabbitMQ()

	ReceiveFromRabbitMQ()
}
