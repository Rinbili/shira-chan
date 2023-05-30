package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Order struct {
	ent.Schema
}

func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Orders"},
		entgql.RelayConnection(),
		entgql.QueryField(),
		//entgql.MultiOrder(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Text("title").
			MaxLen(100).
			Annotations(
				entgql.OrderField("TITLE")).
			Comment("标题").
			NotEmpty(),
		field.Text("content").
			MaxLen(15000).
			//Annotations(
			//	entgql.OrderField("CONTENT")).
			Comment("内容").
			NotEmpty(),
		field.Text("contact").
			MaxLen(20).
			Comment("联系方式"),
		field.Enum("type").
			NamedValues("SOFTWARE", "software",
				"HARDWARE", "hardware",
				"UNKNOWN", "unknown",
				"OTHER", "other").
			Default("other").
			Annotations(
				entgql.OrderField("TYPE")).
			Comment("故障类别"),
		field.Enum("status").
			NamedValues("REQUESTED", "requested",
				"RECEIVED", "received",
				"FINISHED", "finished",
				"CLOSED", "closed").
			Annotations(
				entgql.OrderField("STATUS")).
			Default("requested").
			Comment("工单状态"),
		field.Float("evaluation").
			Annotations(
				entgql.OrderField("EVALUATION")).
			Nillable().
			Optional().
			Comment("工单状态"),
		field.Time("hope_at").
			Default(time.Now()).
			Annotations(
				entgql.OrderField("HOPE_AT")).
			Comment("期望时间").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREAT_AT")).
			Comment("创建时间").
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				entgql.OrderField("UPDATED_AT")).
			Comment("更新时间").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("requester", User.Type).
			Ref("requests").
			Unique().
			Comment("需求者"),
		edge.To("receives", User.Type).
			//Annotations(
			//	entgql.RelayConnection(),
			//	entgql.OrderField("RECEIVES_COUNT")).
			Comment("接单者"),
	}
}

//func (Order) Indexes() []ent.Index {
//	return []ent.Index{
//		index.Fields("id").
//			Edges("requester").
//			Unique(),
//	}
//}
