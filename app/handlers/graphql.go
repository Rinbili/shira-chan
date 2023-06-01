package handlers

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"shira-chan-dev/app/utils"
	"shira-chan-dev/graph"

	"github.com/99designs/gqlgen/graphql/handler"
)

func GraphqlHandler() gin.HandlerFunc {
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewSchema(utils.Client))
	h.Use(entgql.Transactioner{TxOpener: utils.Client})
	return func(c *gin.Context) {
		//保留token解析
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
