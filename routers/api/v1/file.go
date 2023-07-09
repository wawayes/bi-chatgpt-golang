package v1

import (
	"github.com/Walk2future/bi-chatgpt-golang-python/common/requests"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/r"
	"github.com/Walk2future/bi-chatgpt-golang-python/service"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
//	@Router		/gen [post]
func GenChart(c *gin.Context) {
	multipartFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, r.FAIL.WithMsg("获取文件失败"))
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
		c.JSON(http.StatusBadRequest, r.SYSTEM_ERROR.WithMsg(err.Error()))
		log.Println(err.Error())
		return
	}
	// 校验文件
	size := multipartFile.Size
	originalFilename := multipartFile.Filename
	// 校验文件大小
	const ONE_MB = 1024 * 1024
	if size > ONE_MB {
		c.JSON(http.StatusBadRequest, r.FAIL.WithMsg("文件大小超过1M"))
		return
	}
	// 校验文件后缀
	fileSuffix := strings.TrimPrefix(filepath.Ext(originalFilename), ".")
	allowedFileTypes := []string{"xlsx", "xls"}
	if !strutil.ContainsAny(fileSuffix, allowedFileTypes) {
		c.JSON(http.StatusBadRequest, r.FAIL.WithMsg("文件后缀非法"))
		return
	}
	open, err := multipartFile.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, r.SYSTEM_ERROR.WithMsg("文件解析错误"))
		return
	}
	data, err := service.File2Data(open)
	if err != nil {
		c.JSON(http.StatusInternalServerError, r.FAIL.WithMsg("文件读取数据错误"))
		return
	}
	resp := service.GetChatResp(data, goal, chartType)
	c.JSON(http.StatusOK, r.OK.WithData(resp))
}
