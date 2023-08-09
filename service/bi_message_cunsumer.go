package service

import (
	"fmt"
	"strconv"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wawayes/bi-chatgpt-golang/common/constant"
	"github.com/wawayes/bi-chatgpt-golang/models"
	"github.com/wawayes/bi-chatgpt-golang/pkg/logx"
)

type BiMessageConsumer struct {
	channel *amqp.Channel
}

func NewBiMessageConsumer(channel *amqp.Channel) *BiMessageConsumer {
	return &BiMessageConsumer{
		channel: channel,
	}
}

func (c *BiMessageConsumer) Consume() {
	msgs, err := c.channel.Consume(constant.BiQueueName, "", true, false, false, false, nil)
	if err != nil {
		logx.Info(fmt.Sprintf("consumer error %v", err.Error()))
		return
	}
	logx.Info("Consumer执行正常...")

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			chartId, err := strconv.ParseInt(string(msg.Body), 10, 64)
			if err != nil {
				logx.Info("parse chart id error: " + err.Error())
				msg.Nack(false, false)
				continue
			}
			// 利用chartID查询
			chartInfo, err := GetChartById(int(chartId))
			if err != nil {
				return
			}
			if chartInfo == nil {
				logx.Info("图表为空")
				msg.Nack(false, false)
			}
			// 先修改图表任务状态为 “执行中”。等执行成功后，修改为 “已完成”、保存执行结果；执行失败后，状态修改为 “失败”，记录任务失败信息。
			chartInfo.Status = "running"
			logx.Info(fmt.Sprintf("当前正在操作的ChartID: %v", chartInfo.ID))
			err = models.BI_DB.Model(&chartInfo).Updates(chartInfo).Where("id = ?", chartInfo.ID).Error
			if err != nil {
				logx.Info("Chart状态更新失败")
				HandleChartUpdateError(chartInfo.ID, "更新图表执行中状态失败")
				return
			}
			// 调用AI
			content, err := DoChat(BuildUserInput(chartInfo))
			if err != nil {
				logx.Info("调用AI失败")
				return
			}
			// 分隔符
			delimiter := "【【【【【\n"
			parts := strings.Split(content, delimiter)
			if len(parts) < 3 {
				logx.Warning("AI生成结果错误，我最近有种大模型不行了的感觉。。")
				msg.Nack(false, false)
				HandleChartUpdateError(chartInfo.ID, "AI 生成错误")
				return
			}
			genChart := strings.TrimSpace(parts[1])
			genResult := strings.TrimSpace(parts[2])
			chartInfo.GenChart = genChart
			chartInfo.GenResult = genResult
			chartInfo.Status = "succeed"
			if err = models.BI_DB.Model(&chartInfo).Updates(&chartInfo).Where("id = ?", chartInfo.ID).Error; err != nil {
				logx.Info("更新图表状态失败")
				msg.Nack(false, false)
				return
			}
			msg.Ack(false)
		}
	}()
	<-forever
}
