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
	user := apiv1.Group("/user")
	chart := apiv1.Group("/chart")
	table := apiv1.Group("/table")

	user.POST("/login", v1.Login)
	user.POST("/register", v1.Register)
	//apiv1.GET("/current", v1.Current)
	user.Use(auth.MiddlewareFunc())
	{
		user.GET("/refresh_token", v1.RefreshToken)
		user.GET("/list", v1.List)
		user.GET("/current", v1.Current)
		user.GET("/logout", v1.Logout)
	}

	chart.Use(auth.MiddlewareFunc())
	{
		chart.POST("/gen", v1.GenChart)
		chart.POST("/list", v1.ListChart)
		chart.POST("/listALl", v1.ListAllChart)
	}

	table.Use(auth.MiddlewareFunc())
	{
		table.POST("/list", v1.ListUserTable)
	}
	return r
}
