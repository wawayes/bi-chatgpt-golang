package jwt

import (
	"encoding/json"
	"github.com/Walk2future/bi-chatgpt-golang-python/common/requests"
	"github.com/Walk2future/bi-chatgpt-golang-python/models"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/r"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/rsa"
	"github.com/Walk2future/bi-chatgpt-golang-python/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// InitAuth 初始化jwt中间件
func InitAuth() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "bi",                                               // jwt标识
		Key:             []byte("WaWaYes"),                                  // 服务端密钥
		Timeout:         time.Hour * time.Duration(24),                      // token过期时间
		MaxRefresh:      time.Hour * time.Duration(1),                       // token最大刷新时间(RefreshToken过期时间=Timeout+MaxRefresh)
		IdentityHandler: identityHandler,                                    // 解析Claims
		Authenticator:   login,                                              // 校验token的正确性, 处理登录逻辑
		Authorizator:    authorizator,                                       // 用户登录校验成功处理
		Unauthorized:    unauthorized,                                       // 用户登录校验失败处理
		LoginResponse:   loginResponse,                                      // 登录成功后的响应
		LogoutResponse:  logoutResponse,                                     // 登出后的响应
		RefreshResponse: refreshResponse,                                    // 刷新token后的响应
		TokenLookup:     "header: Authorization, query: token, cookie: jwt", // 自动在这几个地方寻找请求中的token
		TokenHeadName:   "Bearer",                                           // header名称
		TimeFunc:        time.Now,
	})
	return authMiddleware, err
}

// 解析Claims
func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	// 此处返回值类型map[string]interface{}与payloadFunc和authorizator的data类型必须一致, 否则会导致授权失败还不容易找到原因
	return map[string]interface{}{
		"IdentityKey": claims[jwt.IdentityKey],
		"user":        claims["user"],
	}
}

// 校验token的正确性, 处理登录逻辑
func login(c *gin.Context) (interface{}, error) {
	var req requests.UserLoginRequest
	// 请求json绑定
	if err := c.ShouldBind(&req); err != nil {
		return "", err
	}

	// 密码通过RSA解密
	decodeData, err := rsa.RSADecrypt([]byte(req.UserPassword), []byte("WaWaYes"))
	if err != nil {
		return nil, err
	}

	r := &requests.UserLoginRequest{
		UserAccount:  req.UserAccount,
		UserPassword: string(decodeData),
	}

	// 密码校验
	var userService service.UserService
	user, err := userService.Login(r)
	if err != nil {
		return nil, err
	}
	// 将用户以json格式写入, payloadFunc/authorizator会使用到
	u, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"user": u,
	}, nil
}

// 用户登录校验成功处理
func authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(map[string]interface{}); ok {
		userStr := v["user"].(string)
		var user models.User
		// 将用户json转为结构体
		err := json.Unmarshal([]byte(userStr), &user)
		if err != nil {
			return false
		}
		// 将用户保存到context, api调用时取数据方便
		c.Set("user", user)
		return true
	}
	return false
}

// 用户登录校验失败处理
func unauthorized(c *gin.Context, code int, message string) {
	r.NO_AUTH.WithMsg("JWT认证失败")
}

// 登录成功后的响应
func loginResponse(c *gin.Context, code int, token string, expires time.Time) {
	c.JSON(http.StatusOK,
		r.OK.WithData(map[string]interface{}{
			"token":   token,
			"expires": expires.Format("2006-01-02 15:04:05"),
		}))
}

// 登出后的响应
func logoutResponse(c *gin.Context, code int) {
	c.JSON(http.StatusOK, r.OK)
}

// 刷新token后的响应
func refreshResponse(c *gin.Context, code int, token string, expires time.Time) {
	c.JSON(http.StatusOK,
		r.OK.WithData(map[string]interface{}{
			"token":   token,
			"expires": expires,
		}))
}
