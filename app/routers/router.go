package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"shira-chan-dev/app/handlers"
)

func InitRouter(r *gin.Engine) {
	//cors
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
	corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	r.Use(cors.New(corsConf))
	//
	r.POST("/query", handlers.AuthHandler(), handlers.GraphqlHandler())
	r.GET("/graphql", handlers.PlaygroundHandler())
}
