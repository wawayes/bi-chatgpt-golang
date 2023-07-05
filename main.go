package main

import (
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/logx"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/setting"
	"github.com/Walk2future/bi-chatgpt-golang-python/routers"
	"log"
	"net/http"
)

//	@title			BI Pro API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8888
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	router := routers.InitRouter()
	go func() {
		log.Println(http.ListenAndServe(setting.Addr, nil))
	}()
	err := router.Run(":8888")
	if err != nil {
		logx.Info("启动成功。。。")
		return
	}
	logx.Error("启动失败。。。")
}
