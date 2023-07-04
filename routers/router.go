package routers

import (
	v1 "github.com/Walk2future/bi-chatgpt-golang-python/routers/api/v1"
	"github.com/gin-gonic/gin"
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

	// 路由
	apiv1 := r.Group("/api/v1")

	apiv1.POST("/login", v1.UserLogin)
	apiv1.POST("/register", v1.UserRegister)

	return r
}
