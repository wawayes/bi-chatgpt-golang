package v1

import (
	"github.com/Walk2future/bi-chatgpt-golang-python/common/requests"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/r"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func GenChart(c *gin.Context) {
	mutipartFile, err := c.FormFile("file")
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
	size := mutipartFile.Size
	originalFilename := mutipartFile.Filename
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
	c.JSON(http.StatusOK, r.OK.WithData(map[string]interface{}{
		"file":      mutipartFile,
		"goal":      goal,
		"chartType": chartType,
	}))
}
