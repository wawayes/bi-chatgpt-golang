package routers

import (
	_ "github.com/Walk2future/bi-chatgpt-golang-python/docs"
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

	// 日志格式化输出到控制台
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	auth, err := jwt.InitAuth()
	if err != nil {
		return nil
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := jwt.AuthMiddleware

	// 路由
	apiv1 := r.Group("/api/v1")
	//apiv1.POST("/login", )
	apiv1.POST("/login", auth.LoginHandler)
	apiv1.GET("/refresh_token", auth.RefreshHandler)

<<<<<<< HEAD
	apiv1.POST("/register", v1.UserRegister)
=======
	apiv1.POST("/login", auth.LoginHandler)
	apiv1.POST("/register", v1.Register)
>>>>>>> origin/dev
	return r
}
