package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
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
		////多字段排序，貌似用不到
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
			////内容排序，应该用不到
			//Annotations(
			//	entgql.OrderField("CONTENT")).
			Comment("内容").
			NotEmpty(),
		field.Text("contact").
			MaxLen(20).
			Comment("联系方式").
			NotEmpty(),
		field.Text("type").
			Default("other").
			Annotations(
				entgql.OrderField("TYPE")).
			Comment("故障类别"),
		field.Bool("is_closed").
			Default(false).
			Annotations(
				entgql.OrderField("IS_CLOSED")).
			Comment("是否被关闭"),
		field.Bool("is_finished").
			Default(false).
			Annotations(
				entgql.OrderField("IS_FINISHED")).
			Comment("是否完成"),
		field.Float("evaluation").
			Annotations(
				entgql.OrderField("EVALUATION")).
			Optional().
			Min(1).
			Max(5).
			Comment("评分").
			Nillable(),
		field.Int64("hope_at").
			DefaultFunc(time.Now().Unix).
			Annotations(
				entgql.OrderField("HOPE_AT"),
				entgql.Type("Int64")).
			Comment("期望时间"),
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("requester", User.Type).
			Ref("requested").
			Unique().
			Comment("需求者"),
		edge.To("receiver", User.Type).
			Annotations(
				entgql.RelayConnection(),
				entgql.OrderField("RECEIVER_COUNT")).
			Comment("接单者"),
	}
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
