package main

import (
	"github.com/robfig/cron"
	"github.com/wawayes/bi-chatgpt-golang/pkg/logx"
	"github.com/wawayes/bi-chatgpt-golang/routers"
	"log"
	"time"
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
	//go func() {
	//	log.Println(http.ListenAndServe("http://localhost", nil))
	//}()
	err := router.Run(":8888")
	if err != nil {
		logx.Info("启动成功。。。")
		return
	}
	logx.Error("启动失败。。。")
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		//models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		//models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
