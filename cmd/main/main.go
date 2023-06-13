package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"shira-chan-dev/app/routers"
)

func main() {
	r := gin.Default()
	// 初始化路由
	routers.InitRouter(r)
	// 运行
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
