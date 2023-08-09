package service

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/pandodao/tokenizer-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wawayes/bi-chatgpt-golang/common/constant"
	"github.com/wawayes/bi-chatgpt-golang/common/requests"
	"github.com/wawayes/bi-chatgpt-golang/common/response"
	"github.com/wawayes/bi-chatgpt-golang/models"
	"github.com/wawayes/bi-chatgpt-golang/pkg/logx"
	"github.com/xuri/excelize/v2"
)

// Xlsx2Data 读取xlsx文件数据
func Xlsx2Data(file multipart.File) (data string, err error) {
	// TODO 将Xlsx转为CSV文件，经过实际测试，token数并不会有什么变化
	f, err := excelize.OpenReader(file)
	if err != nil {
		logx.Warning(err.Error())
		return "", err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for _, row := range rows {
		for _, colCell := range row {
			data += colCell + "\t"
		}
		data += "\n"
	}
	return data, nil
}

// GetChatResp 获取ChatGPT响应
func GetChatResp(c *gin.Context, newChart models.Chart) (res response.BiResp, err error) {
	// 建立连接
	conn, err := amqp.Dial(constant.MQUrl)
	if err != nil {
		fmt.Printf("dial rabbitmq error: %s\n", err)
		return
	}
	defer conn.Close()

	// 创建channel
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("create channel error: %s\n", err)
		return
	}
	defer ch.Close()

	// 声明交换机
	if err = ch.ExchangeDeclare(
		constant.BiExchangeName, // name
		"direct",                // type
		true,                    // durable
		false,                   // auto-deleted
		false,                   // internal
		false,                   // no-wait
		nil,                     // arguments
	); err != nil {
		fmt.Printf("declare exchange error: %s\n", err)
		return
	}

	// 声明队列
	q, err := ch.QueueDeclare(
		constant.BiQueueName, // name
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		fmt.Printf("declare queue error: %s\n", err)
		return
	}

	// 绑定队列到交换机
	if err = ch.QueueBind(
		q.Name,                  // queue name
		constant.BiRoutingKey,   // routing key
		constant.BiExchangeName, // exchange
		false,
		nil); err != nil {
		fmt.Printf("bind queue error: %s\n", err)
		return
	}

	// 创建消费者
	consumer := NewBiMessageConsumer(ch)
	logx.Info(fmt.Sprintf("Channel信息: %v", ch))
	go consumer.Consume()

	prompt := BuildUserInput(&newChart)
	content, err := DoChat(prompt)
	if err != nil {
		return response.BiResp{}, err
	}
	var biResp response.BiResp
	//var userService *UserService
	// 计算token值
	OriginStr := SystemPrompt + prompt + content
	t := tokenizer.MustCalToken(OriginStr)
	userService := &UserService{}
	current, _ := userService.Current(c)
	newChart.UserId = current.ID
	newChart.Token = t
	var user models.User
	if err = models.BI_DB.Model(&user).Where("id = ?", current.ID).First(&user).Error; err != nil {
		return response.BiResp{}, errors.New("查找当前用户失败")
	}
	user.FreeCount--
	if err := models.BI_DB.Save(&user).Error; err != nil {
		return response.BiResp{}, errors.New("FreeCount--异常")
	}
	var chartInfo *models.Chart
	models.BI_DB.Model(&chartInfo).Select("*").Where("id = ?", newChart.ID).First(&chartInfo)
	biResp.GenResult = chartInfo.GenResult
	biResp.GenChart = chartInfo.GenChart
	return biResp, nil
}

// GetChartById 根据id获取Chart
func GetChartById(chartId int) (*models.Chart, error) {
	var chartInfo models.Chart
	if err := models.BI_DB.Model(&chartInfo).Select("*").Where("id = ?", chartId).First(&chartInfo).Error; err != nil {
		logx.Info("根据Id查询Chart失败: " + err.Error())
		return nil, err
	}
	return &chartInfo, nil
}

// ListChart 分页查询当前用户图表
func ListChart(c *gin.Context, chartQueryRequest *requests.ChartQueryRequest) ([]models.Chart, error) {
	var userService UserService
	currentUser, err := userService.Current(c)
	if err != nil {
		return nil, errors.New("获取当前用户失败")
	}
	userId := currentUser.ID
	chartQueryRequest.UserId = userId
	pageNum := chartQueryRequest.PageNum
	pageSize := chartQueryRequest.PageSize
	if pageSize > 20 {
		return nil, errors.New("你要的页数太多了")
	}
	var chartList []models.Chart
	if err := models.BI_DB.Model(&chartList).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&chartList).Error; err != nil {
		return nil, errors.New("分页查询当前用户图表失败")
	}
	return chartList, nil
}

// ListAllChart 分页查询所有用户图表
func ListAllChart(chartQueryRequest *requests.ChartQueryRequest) (listAllChart []models.Chart, err error) {
	pageNum := chartQueryRequest.PageNum
	pageSize := chartQueryRequest.PageSize
	if err = models.BI_DB.Model(&listAllChart).Select("").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&listAllChart).Error; err != nil {
		return nil, errors.New("数据库查询listAllChart失败")
	}
	return listAllChart, nil
}
