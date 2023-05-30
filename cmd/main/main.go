package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"shira-chan-dev/app/handlers"
)

func main() {
	r := gin.Default()
	r.POST("/query", handlers.GraphqlHandler())
	r.GET("/graphql", handlers.PlaygroundHandler())
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
