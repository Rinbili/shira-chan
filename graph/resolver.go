package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"shira-chan-dev/ent"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers:  &Resolver{client},
		Directives: DirectiveRoot{},
	})
}
