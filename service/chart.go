package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Walk2future/bi-chatgpt-golang-python/common/requests"
	"github.com/Walk2future/bi-chatgpt-golang-python/common/response"
	"github.com/Walk2future/bi-chatgpt-golang-python/models"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/logx"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// Xlsx2Data 读取xlsx文件数据
func Xlsx2Data(file multipart.File) (data string, err error) {
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
func GetChatResp(c *gin.Context, info string, goal string, chartType string) (res response.BiResp, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	var chatReq requests.ChatRequest
	systemPrompt := "你是一个高级数据分析师和前端开发专家，接下来我按照以下格式给你提供内容：" +
		"\n分析需求：{分析需求和目标}\n原始数据：{原始数据}\nEcharts图表类型：{Echarts图表类型}" +
		"\n请根据这两部分内容按照以下指定格式生成内容（不要输出任何多余的开头或者结尾或者注释）" +
		"\n【【【【【\n{前端的Echarts V5的option配置对象json代码，合理地将数据进行可视化，不要生成多余的开头结尾或者任何注释}" +
		"\n【【【【【\n{明确的数据结论，越详细越好，不要生成任何多余废话或者对实质结论无用的内容}"
	prompt := "原始数据：" + info + "\n分析需求和目标：" + goal + ", Echarts图表类型：" + chartType
	chatReq.Model = "gpt-3.5-turbo"
	chatReq.Messages = []requests.Message{
		{Role: "system", Content: systemPrompt},
		{Role: "user", Content: prompt},
	}
	data, err := json.Marshal(chatReq)
	if err != nil {
		return response.BiResp{}, err
	}
	req, err := http.NewRequest("POST", os.Getenv("BASE_URL"), bytes.NewBuffer(data))
	if err != nil {
		return response.BiResp{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return response.BiResp{}, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return response.BiResp{}, err
	}
	var chatResp response.ChatCompletionResponse
	err = json.Unmarshal(respBody, &chatResp)
	if err != nil {
		return response.BiResp{}, err
	}
	content := chatResp.Choices[0].Message.Content
	var biResp response.BiResp
	delimiter := "【【【【【\n"
	parts := strings.Split(content, delimiter)
	if len(parts) < 3 {
		logx.Warning("AI生成结果错误，我最近有种大模型不行了的感觉。。")
		return response.BiResp{}, err
	}
	for i, part := range parts {
		if i == 1 {
			biResp.GenChart = part
		}
		if i == 2 {
			biResp.GenResult = part
		}
	}
	//var userService *UserService
	userService := &UserService{}
	current, _ := userService.Current(c)
	chart := &models.Chart{
		UserId:    current.ID,
		Data:      info,
		Goal:      goal,
		ChartType: chartType,
		GenChart:  biResp.GenChart,
		GenResult: biResp.GenResult,
	}
	err = models.BI_DB.Model(&models.Chart{}).Select("goal", "chartType", "genChart", "genResult", "userId").Create(&chart).Error
	if err != nil {
		logx.Warning(err.Error())
		return response.BiResp{}, err
	}
	var user models.User
	if err := models.BI_DB.Model(&user).Where("id = ?", current.ID).First(&user).Error; err != nil {
		return response.BiResp{}, errors.New("查找当前用户失败")
	}
	user.FreeCount--
	if err := models.BI_DB.Save(&user).Error; err != nil {
		return response.BiResp{}, errors.New("FreeCount--异常")
	}
	return biResp, nil
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
