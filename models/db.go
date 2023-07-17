package models

import (
	"fmt"
	"github.com/wawayes/bi-chatgpt-golang/pkg/logx"
	"github.com/wawayes/bi-chatgpt-golang/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var BI_DB *gorm.DB

func init() {
	var (
		err                              error
		dbUser, dbPassword, path, dbName string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		logx.Warning(fmt.Sprintf("获取数据库配置失败:%v", err.Error()))
		return
	}
	dbUser = sec.Key("USER").String()
	dbPassword = sec.Key("PASSWORD").String()
	path = sec.Key("HOST").String()
	dbName = sec.Key("NAME").String()

	BI_DB, err = gorm.Open(mysql.New(mysql.Config{
		//DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, path, dbName),
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logx.Error(err.Error())
	}
	fmt.Println("数据库连接成功")

}
