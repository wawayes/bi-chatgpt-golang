package main

import (
	"fmt"
	"github.com/Walk2future/bi-chatgpt-golang-python/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:8000", nil))
	}()
	err := router.Run()
	if err != nil {
		fmt.Println("启动失败...")
		return
	}
	fmt.Println("启动成功...")
}
