package v1

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/wawayes/bi-chatgpt-golang/common/constant"
	"github.com/wawayes/bi-chatgpt-golang/common/requests"
	"github.com/wawayes/bi-chatgpt-golang/models"
	"github.com/wawayes/bi-chatgpt-golang/pkg/logx"
	"github.com/wawayes/bi-chatgpt-golang/pkg/r"
	"github.com/wawayes/bi-chatgpt-golang/service"
)

// GenChart godoc
//
//	@Summary		生成图表
//	@Description	通过上传文件和发送请求生成图表
//	@Accept			multipart/form-data
//	@Produce		json
//	@Tags			ChartApi
//	@Param			file		formData	file			true	"要上传的文件"
//	@Param			goal		formData	string			true	"生成图表的目标"
//	@Param			chartType	formData	string			true	"图表类型"
//	@Success		200			{object}	response.BiResp	"成功"
//	@Failure		40002		{object}	r.Response		"参数错误"
//	@Failure		40003		{object}	r.Response		"系统错误"
//	@Router			/chart/gen [post]
func GenChart(c *gin.Context) {
	// 建立连接
	conn, err := amqp.Dial(constant.MQUrl)
	if err != nil {
		fmt.Printf("dial rabbitmq error: %s\n", err)
		return
	}
	logx.Info("连接MQ成功")
	defer conn.Close()

	// 创建channel
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("create channel error: %s\n", err)
		return
	}
	logx.Info("创建channel成功")
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
	logx.Info("声明交换机成功")
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
	logx.Info("声明队列成功")
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

	multipartFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, r.FAIL.WithMsg("获取文件失败"))
		return
	}
	goal := c.Request.FormValue("goal")
	chartType := c.Request.FormValue("chartType")
	req := requests.GenRequest{
		Goal:      goal,
		ChartType: chartType,
	}
	validate := validator.New()
	// 使用validator库进行参数校验
	if err := validate.Struct(&req); err != nil {
		c.JSON(http.StatusOK, r.SYSTEM_ERROR.WithMsg(err.Error()))
		log.Println(err.Error())
		return
	}
	// 校验文件
	size := multipartFile.Size
	originalFilename := multipartFile.Filename
	// 校验文件大小
	const OneMb = 1024 * 1024
	if size > OneMb {
		c.JSON(http.StatusOK, r.FAIL.WithMsg("文件大小超过1M"))
		return
	}
	// 校验文件后缀
	fileSuffix := strings.TrimPrefix(filepath.Ext(originalFilename), ".")
	allowedFileTypes := []string{"xlsx", "xls"}
	if !strutil.ContainsAny(fileSuffix, allowedFileTypes) {
		c.JSON(http.StatusOK, r.FAIL.WithMsg("文件后缀非法"))
		return
	}

	open, err := multipartFile.Open()
	if err != nil {
		c.JSON(http.StatusOK, r.SYSTEM_ERROR.WithMsg("文件解析错误"))
		return
	}
	data, err := service.Xlsx2Data(open)
	if err != nil {
		c.JSON(http.StatusOK, r.FAIL.WithMsg("文件读取数据错误"))
		return
	}
	var userService service.UserService
	currentUser, err := userService.Current(c)
	if err != nil {
		logx.Info("获取当前用户失败")
	}
	newChart := models.Chart{
		Goal:      goal,
		Data:      data,
		ChartType: chartType,
		Status:    "wait",
		UserId:    currentUser.ID,
	}
	if err = models.BI_DB.Model(&newChart).Create(&newChart).Error; err != nil {
		logx.Warning("保存newChart失败: " + err.Error())
	}
	producer := service.NewBiMessageProducer(ch)
	err = producer.Publish(c, strconv.FormatInt(int64(newChart.ID), 10))
	if err != nil {
		logx.Warning("发布消息失败：" + err.Error())
		c.JSON(http.StatusOK, r.SYSTEM_ERROR)
	}
	res, err := service.GetChatResp(c, newChart)
	if err != nil {
		logx.Warning("生成响应失败" + err.Error())
	}
	c.JSON(http.StatusOK, r.OK.WithData(res))
}

// ListChart godoc
//
//	@Summary	Chart List
//	@Produce	json
//	@Tags		ChartApi
//	@Param		ChartQueryRequest	body	requests.ChartQueryRequest	true	"查询请求参数"
//	@Accept		multipart/form-data
//	@Success	0		{object}	response.BiResp	"成功"
//	@Failure	40002	{object}	r.Response		"参数错误"
//	@Failure	40003	{object}	r.Response		"系统错误"
//	@Router		/chart/list [post]
func ListChart(c *gin.Context) {
	var req requests.ChartQueryRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, r.PARAMS_ERROR.WithMsg("参数有误"))
		return
	}
	listChart, err := service.ListChart(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, r.FAIL.WithMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, r.OK.WithData(listChart))
}

// ListAllChart godoc
//
//	@Summary	ListAllChart
//	@Produce	json
//	@Tags		ChartApi
//	@Param		ChartQueryRequest	body	requests.ChartQueryRequest	true	"查询请求参数"
//	@Accept		multipart/form-data
//	@Success	0		{object}	response.BiResp	"成功"
//	@Failure	40002	{object}	r.Response		"参数错误"
//	@Failure	40003	{object}	r.Response		"系统错误"
//	@Router		/chart/all_list [post]
func ListAllChart(c *gin.Context) {
	var req requests.ChartQueryRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, r.PARAMS_ERROR.WithMsg("参数错误"))
		return
	}
	allChart, err := service.ListAllChart(&req)
	if err != nil {
		c.JSON(http.StatusOK, r.FAIL.WithMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, r.OK.WithData(allChart))
}
