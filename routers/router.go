package routers

import (
	_ "github.com/Walk2future/bi-chatgpt-golang-python/docs"
	"github.com/Walk2future/bi-chatgpt-golang-python/middleware/cors"
	"github.com/Walk2future/bi-chatgpt-golang-python/middleware/jwt"
	v1 "github.com/Walk2future/bi-chatgpt-golang-python/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func InitRouter() *gin.Engine {
	// 强制日志高亮
	gin.ForceConsoleColor()
	r := gin.Default()
	r.Use(cors.Cors())
	//r.Use(func(c *gin.Context) {
	//	// 设置cookie的过期时间为一个较早的时间点，比如1970年1月1日
	//	c.SetCookie("jwt", "", -1, "/", "localhost", false, true)
	//	c.Next()
	//})

	// 日志格式化输出到控制台
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	var auth = *jwt.AuthMiddleware

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 路由
	apiv1 := r.Group("/api/v1")

	apiv1.POST("/login", v1.Login)
	apiv1.GET("/refresh_token", v1.RefreshToken)
	//apiv1.GET("/current", v1.Current)
	apiv1.Use(auth.MiddlewareFunc())
	{
		apiv1.GET("/current", v1.Current)
		apiv1.GET("/logout", v1.Logout)
	}
	apiv1.POST("/register", v1.Register)
	return r
}
