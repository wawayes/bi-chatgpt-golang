package redis

import (
	"fmt"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/logx"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func init() {
	var (
		err                                            error
		redisPath, redisAddr, redisPort, redisPassword string
		redisDB                                        int
	)
	sec, err := setting.Cfg.GetSection("redis")
	if err != nil {
		logx.Warning(fmt.Sprintf("获取数据库配置失败:%v", err.Error()))
		return
	}
	redisPath = sec.Key("PATH").String()
	redisPort = sec.Key("PORT").String()
	redisAddr = fmt.Sprintf(redisPath + ":" + redisPort)
	redisPassword = fmt.Sprintf(sec.Key("PASSWORD").MustString(""))
	redisDB = sec.Key("DB").MustInt(0)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword, // no password set
		DB:       redisDB,       // use default DB
	})

}