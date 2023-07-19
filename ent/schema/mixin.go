package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type TimeMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("created_at").
			DefaultFunc(func() int64 { return time.Now().Unix() }).
			Annotations(
				entgql.OrderField("CREAT_AT"),
				entgql.Type("Int64")).
			Comment("创建时间").
			Immutable(),
		field.Int64("updated_at").
			DefaultFunc(func() int64 { return time.Now().Unix() }).
			UpdateDefault(func() int64 { return time.Now().Unix() }).
			Annotations(
				entgql.OrderField("UPDATED_AT"),
				entgql.Type("Int64")).
			Comment("更新时间"),
	}
}
