package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"shira-chan-dev/app/routers"
	"shira-chan-dev/config"
)

func main() {
	r := gin.Default()
	// 初始化路由
	routers.InitRouter(r)
	// 运行
	if config.Config.Ssl.Enable {
		err := r.RunTLS(":"+config.Config.Port, config.Config.Ssl.Crt, config.Config.Ssl.Key)
		if err != nil {
			panic(err)
		}
	} else {
		err := r.Run(":" + config.Config.Port)
		if err != nil {
			panic(err)
		}
	}
}
