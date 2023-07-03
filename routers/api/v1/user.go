package v1

import (
	"github.com/Walk2future/bi-chatgpt-golang-python/common/requests"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/logx"
	"github.com/Walk2future/bi-chatgpt-golang-python/service"
	"log"

	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/r"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func UserLogin(c *gin.Context) {
	userService := service.UserService{}
	var req requests.UserLoginRequest
	validate := validator.New()
	// 使用validator库进行参数校验
	if err := validate.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, r.SYSTEM_ERROR.WithMsg(err.Error()))
		log.Println(err.Error())
		return
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, r.PARAMS_ERROR.WithMsg("请求参数错误"))
		log.Println(err.Error())
		return
	}
	user, err := userService.UserLogin(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, r.SYSTEM_ERROR.WithMsg("登录失败:"+err.Error()))
	} else {
		c.JSON(http.StatusOK, r.OK.WithData(user))
		logx.Info("用户登录成功")
	}
}
