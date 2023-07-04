package main

import (
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/logx"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/setting"
	"github.com/Walk2future/bi-chatgpt-golang-python/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	go func() {
		log.Println(http.ListenAndServe("localhost:"+setting.HTTPPort, nil))
	}()
	err := router.Run()
	if err != nil {
		logx.Info("启动成功。。。")
		return
	}
	logx.Error("启动失败。。。")
}
