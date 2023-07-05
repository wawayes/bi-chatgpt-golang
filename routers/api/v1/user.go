package v1

import (
	"context"
	"fmt"
	"github.com/Walk2future/bi-chatgpt-golang-python/common/requests"
	"github.com/Walk2future/bi-chatgpt-golang-python/middleware/redis"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/logx"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/r"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/session"
	"github.com/Walk2future/bi-chatgpt-golang-python/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	_ "github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

// UserLogin godoc
//
//	@Summary	User Login
//	@Produce	json
//	@Tags		UserApi
//	@Param		userAccount		body	string	true	"用户名"
//	@Param		userPassword	body	string	true	"密码"
//	@Accept		json
//	@Success	0		{object}	models.User	"成功"
//	@Failure	40002	{object}	r.Response	"参数错误"
//	@Failure	40003	{object}	r.Response	"系统错误"
//	@Router		/login [post]
func UserLogin(c *gin.Context) {
	userService := &service.UserService{}
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
		if err != nil {
			logx.Error("user信息解码失败")
			panic(err)
		}
		s := &session.Session{
			SessionID: uuid.New(),
			UserInfo:  *user,
		}
		logx.Info("用户登录成功")
		err = s.Save(context.Background(), redis.Rdb)
		if err != nil {
			logx.Warning("登录信息存入session失败")
			return
		}
		logx.Info(fmt.Sprintf("登录信息存入session成功~!:%v", s))
	}
}

// UserRegister godoc
//
//	@Summary	User Register
//	@Produce	json
//	@Tags		UserApi
//	@Param		userAccount		body	string	true	"用户名"
//	@Param		userPassword	body	string	true	"密码"
//	@Param		checkPassword	body	string	true	"检查密码"
//	@Accept		json
//	@Success	0		{object}	models.User	"成功"
//	@Failure	40002	{object}	r.Response	"参数错误"
//	@Failure	40003	{object}	r.Response	"系统错误"
//	@Router		/register [post]
func UserRegister(c *gin.Context) {
	userService := &service.UserService{}
	var req requests.UserRegisterRequest
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
	res, err := userService.UserRegister(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, r.SYSTEM_ERROR.WithMsg("注册失败:"+err.Error()))
	} else {
		c.JSON(http.StatusOK, r.OK.WithData(res))
	}
}
