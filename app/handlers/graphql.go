package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"shira-chan-dev/ent"
	"shira-chan-dev/ent/migrate"
	"shira-chan-dev/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

// Defining the Graphql handler
func GraphqlHandler() gin.HandlerFunc {
	// Create ent.Client and run the schema migration.
	dsn := os.Getenv("sql_user") + ":" +
		os.Getenv("sql_passwd") + "@tcp(" +
		os.Getenv("sql_host") + ":" +
		os.Getenv("sql_port") + ")/" +
		os.Getenv("sql_name") + "?charset=utf8mb4&parseTime=True&loc=Local"
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewSchema(client))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
