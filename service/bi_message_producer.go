package service

import (
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wawayes/bi-chatgpt-golang/common/constant"
	"github.com/wawayes/bi-chatgpt-golang/pkg/logx"
)

// BiMessageProducer 生产者
type BiMessageProducer struct {
	channel *amqp.Channel
}

// NewBiMessageProducer 创建生产者
func NewBiMessageProducer(channel *amqp.Channel) *BiMessageProducer {
	return &BiMessageProducer{
		channel: channel,
	}
}

// Publish 发布消息
func (p *BiMessageProducer) Publish(ctx *gin.Context, message string) error {
	logx.Info("Producer生产者发布消息...")
	return p.channel.PublishWithContext(
		ctx,
		constant.BiExchangeName, // exchange
		constant.BiRoutingKey,   // routing key
		false,                   // mandatory
		false,                   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}
