package v1

import (
	"github.com/Walk2future/bi-chatgpt-golang-python/common/requests"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/r"
	"github.com/Walk2future/bi-chatgpt-golang-python/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

// Login godoc
//
//	@Summary	User Login
//	@Produce	json
//	@Tags		UserApi
//	@Param		loginRequest	body	requests.UserLoginRequest	true	"登录请求参数"
//	@Accept		json
//	@Success	0		{object}	session.Session	"成功"
//	@Failure	40002	{object}	r.Response		"参数错误"
//	@Failure	40003	{object}	r.Response		"系统错误"
//	@Router		/login [post]
<<<<<<< HEAD
//func UserLogin(c *gin.Context) {
//	userService := &service.UserService{}
//	var req requests.LoginRequest
//	validate := validator.New()
//	// 使用validator库进行参数校验
//	if err := validate.Struct(&req); err != nil {
//		c.JSON(http.StatusBadRequest, r.SYSTEM_ERROR.WithMsg(err.Error()))
//		log.Println(err.Error())
//		return
//	}
//	if err := c.BindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, r.PARAMS_ERROR.WithMsg("请求参数错误"))
//		log.Println(err.Error())
//		return
//	}
//	user, err := userService.UserLogin(&req)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, r.SYSTEM_ERROR.WithMsg("登录失败:"+err.Error()))
//	} else {
//		c.JSON(http.StatusOK, r.OK.WithData(user))
//		if err != nil {
//			logx.Error("user信息解码失败")
//			panic(err)
//		}
//		s := &session.Session{
//			SessionID: uuid.New(),
//			UserInfo:  *user,
//		}
//		logx.Info("用户登录成功")
//		err = s.Save(context.Background(), redis.Rdb)
//		if err != nil {
//			logx.Warning("登录信息存入session失败")
//			return
//		}
//		logx.Info(fmt.Sprintf("登录信息存入session成功~!:%v", s))
//	}
//}
=======
func Login(c *gin.Context) {
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
	user, err := userService.Login(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, r.SYSTEM_ERROR.WithMsg("登录失败:"+err.Error()))
	} else {
		// 将用户信息存入session
		s := &session.Session{
			SessionID: uuid.New(),
			UserInfo:  *user,
		}
		err = s.Save(context.Background(), redis.Rdb)
		if err != nil {
			logx.Warning("登录信息存入session失败")
			return
		}
		logx.Info(fmt.Sprintf("登录信息存入session成功~!:%v", s))
		c.JSON(http.StatusOK, r.OK.WithData(s))
		if err != nil {
			logx.Error("user信息解码失败")
			panic(err)
		}
		logx.Info("用户登录成功")
	}
}
>>>>>>> origin/dev

// Register godoc
//
//	@Summary	User Register
//	@Produce	json
//	@Tags		UserApi
//	@Param		registerRequest	body	requests.UserRegisterRequest	true	"注册请求参数"
//	@Accept		json
//	@Success	0		{object}	models.User	"成功"
//	@Failure	40002	{object}	r.Response	"参数错误"
//	@Failure	40003	{object}	r.Response	"系统错误"
//	@Router		/register [post]
func Register(c *gin.Context) {
	userService := &service.UserService{}
	var req requests.RegisterRequest
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
	res, err := userService.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, r.SYSTEM_ERROR.WithMsg("注册失败:"+err.Error()))
	} else {
		c.JSON(http.StatusOK, r.OK.WithData(res))
	}
}

func Current(c *gin.Context) {

}
