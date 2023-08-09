package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/wawayes/bi-chatgpt-golang/common/requests"
	"github.com/wawayes/bi-chatgpt-golang/common/response"
	"github.com/wawayes/bi-chatgpt-golang/models"
	"github.com/wawayes/bi-chatgpt-golang/pkg/logx"
)

var SystemPrompt = "你是一个高级数据分析师和前端开发专家，接下来我按照以下格式给你提供内容：" +
	"\n分析需求：{分析需求和目标}\n原始数据：{原始数据}\nEcharts图表类型：{Echarts图表类型}" +
	"\n请根据这两部分内容按照以下指定格式生成内容（不要输出任何多余的开头或者结尾或者注释）" +
	"\n【【【【【\n{前端的Echarts V5的option配置对象json代码，合理地将数据进行可视化，不要生成多余的开头结尾或者任何注释}" +
	"\n【【【【【\n{明确的数据结论，越详细越好，不要生成任何多余废话或者对实质结论无用的内容}"

// DoChat 获取ChatGPT响应
func DoChat(prompt string) (content string, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	var chatReq requests.ChatRequest
	chatReq.Model = "gpt-3.5-turbo"
	chatReq.Messages = []requests.Message{
		{Role: "system", Content: SystemPrompt},
		{Role: "user", Content: prompt},
	}
	data, err := json.Marshal(chatReq)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", os.Getenv("BASE_URL"), bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var chatResp response.ChatCompletionResponse
	err = json.Unmarshal(respBody, &chatResp)
	if err != nil {
		return "", err
	}
	content = chatResp.Choices[0].Message.Content
	if len(content) == 0 {
		logx.Info("ChatGPT Response为空")
		return "", errors.New("ChatGPT响应为空")
	}
	logx.Info("ChatGPT Response: " + content)
	return content, nil
}

// 构建用户输入
func BuildUserInput(chart *models.Chart) string {
	// 构造用户输入
	var userInput strings.Builder
	userInput.WriteString("分析需求:\n")
	// 拼接分析目标
	userGoal := chart.Goal
	if chart.ChartType != "" {
		userGoal += ",请使用" + chart.ChartType
	}
	userInput.WriteString(userGoal + "\n")
	userInput.WriteString("原始数据:\n")
	userInput.WriteString(chart.Data + "\n")
	return userInput.String()
}

func HandleChartUpdateError(chartId int, execMessage string) {
	var updateChartResult *models.Chart
	updateChartResult.ID = chartId
	updateChartResult.Status = "failed"
	updateChartResult.ExecMessage = "execMessage"
	if err := models.BI_DB.Model(&updateChartResult).Updates(updateChartResult).Where("id = ?", updateChartResult.ID).Error; err != nil {
		logx.Info(fmt.Sprintf("%v更新图表状态失败: %v", updateChartResult.ID, err.Error()))
		return
	}
}
