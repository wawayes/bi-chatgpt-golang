package v1

import (
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wawayes/bi-chatgpt-golang/common/requests"
	"github.com/wawayes/bi-chatgpt-golang/pkg/r"
	"github.com/wawayes/bi-chatgpt-golang/service"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// GenChart godoc
//
//	@Summary	Generate a chart
//	@Produce	json
//	@Tags		ChartApi
//	@Param		file		formData	file				true	"登录请求参数"
//	@Param		genRequest	formData	requests.GenRequest	true	"生成请求"
//	@Accept		multipart/form-data
//	@Success	0		{object}	response.BiResp	"成功"
//	@Failure	40002	{object}	r.Response		"参数错误"
//	@Failure	40003	{object}	r.Response		"系统错误"
//	@Router		/chart/gen [post]
func GenChart(c *gin.Context) {
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
	res, err := service.GetChatResp(c, data, goal, chartType)
	if err != nil || strutil.IsBlank(res.GenChart) || strutil.IsBlank(res.GenResult) {
		c.JSON(http.StatusOK, r.FAIL.WithMsg("我总感觉大模型越来越傻了,别生气,要不再试一次"))
		return
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
