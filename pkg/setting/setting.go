package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	err error
	Cfg *ini.File

	Url      string
	HTTPPort string
)

func init() {
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		fmt.Printf("Failed to load file:%v", err)
		os.Exit(1)
	}
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		return
	}
	Url = sec.Key("url").String()
	HTTPPort = sec.Key("http_port").String()
}
