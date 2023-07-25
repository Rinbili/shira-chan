package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Receive holds the edge schema definition for the Receive edge.
type Receive struct {
	ent.Schema
}

func (Receive) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("order_id", "user_id"),
		entsql.Annotation{Table: "order_receiver"},
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

// Fields of the Receive.
func (Receive) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("order_id"),
	}
}

// Edges of the Receive.
func (Receive) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id").
			Annotations(
				entgql.RelayConnection(),
				entgql.OrderField("RECEIVED_COUNT")),
		edge.To("order", Order.Type).
			Unique().
			Required().
			Field("order_id").
			Annotations(
				entgql.RelayConnection(),
				entgql.OrderField("RECEIVER_COUNT")),
	}
}

func (Receive) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
