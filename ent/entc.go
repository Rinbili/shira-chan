//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./ent.graphql"),
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithWhereInputs(true),
	)
	opts := []entc.Option{
		entc.Extensions(ex),
		entc.FeatureNames("privacy", "schema/snapshot", "entql"),
	}
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	if err := entc.Generate("./../ent/schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
