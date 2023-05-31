package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"shira-chan-dev/app/routers"
)

func main() {
	r := gin.Default()
	routers.InitRouter(r)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
