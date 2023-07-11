package jwt

import (
	"github.com/Walk2future/bi-chatgpt-golang-python/common/requests"
	"github.com/Walk2future/bi-chatgpt-golang-python/models/serializers"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/r"
	"github.com/Walk2future/bi-chatgpt-golang-python/service"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"

var AuthMiddleware *jwt.GinJWTMiddleware

func init() {

	var err error

	// the jwt middleware
	AuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "user info",
		Key:         []byte("BI"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		SendCookie:  true,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*serializers.CurrentUser); ok {
				return jwt.MapClaims{
					identityKey:   v.ID,
					"userAccount": v.UserAccount,
					"userName":    v.UserName,
					"userAvatar":  v.UserAvatar,
					"userRole":    v.UserRole,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &serializers.CurrentUser{
				ID:          claims[identityKey].(string),
				UserAccount: claims["userAccount"].(string),
				UserName:    claims["userName"].(string),
				UserAvatar:  claims["userAvatar"].(string),
				UserRole:    claims["userRole"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals requests.LoginRequest
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userService := service.UserService{}
			user, err := userService.Login(&loginVals)
			if err != nil {
				c.JSON(http.StatusBadRequest, r.PARAMS_ERROR.WithMsg(err.Error()))
				panic(err.Error())
				return nil, err
			}
			// 生成token
			//token, expire, err := AuthMiddleware.TokenGenerator(user)
			//if err != nil {
			//	return nil, err
			//}
			// 将token存入redis
			// 将time.Time转化为Duration
			//duration := time.Until(expire)
			//redisDuration := duration.Round(time.Second)
			//statusCmd := redis.Rdb.Set(context.Background(), "token:"+user.ID, token, redisDuration)
			//if statusCmd.Err() != nil {
			//	return nil, statusCmd.Err()
			//}
			return user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*serializers.CurrentUser); ok && v.UserRole != "banned" {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(http.StatusBadRequest, r.NO_AUTH.WithMsg("认证失败"))
			c.Abort()
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,

		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, r.OK.WithData(map[string]interface{}{
				"token":  token,
				"expire": expire,
			}))
		},
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, r.OK.WithMsg("退出登录成功"))
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, r.OK.WithData(map[string]interface{}{
				"token":  token,
				"expire": expire,
			}))
		},
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := AuthMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

}
