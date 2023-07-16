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

type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Users"},
		entgql.RelayConnection(),
		entgql.QueryField(),
		//entgql.MultiOrder(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Text("uname").
			MaxLen(30).
			Annotations(
				entgql.OrderField("UNAME")).
			Comment("用户名").
			NotEmpty(),
		field.Text("passwd").
			Sensitive().
			Comment("密码").
			NotEmpty(),
		field.Text("phone").
			MinLen(5).
			MaxLen(15).
			Unique().
			Comment("手机号码"),
		field.Bool("is_admin").
			Default(false).
			Annotations(
				entgql.OrderField("IS_ADMIN")).
			Comment("是否管理员"),
		field.Bool("is_secretary").
			Default(false).
			Annotations(
				entgql.OrderField("IS_SECRETARY")).
			Comment("是否部员"),
		field.Bool("is_active").
			Default(true).
			Annotations(
				entgql.OrderField("STATE")).
			Comment("用户状态"),
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

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("requested", Order.Type).
			Annotations(
				entgql.RelayConnection(),
				entgql.OrderField("REQUESTED_COUNT")).
			Comment("需求"),
		edge.From("received", Order.Type).
			Ref("receiver").
			Annotations(
				entgql.RelayConnection(),
				entgql.OrderField("RECEIVED_COUNT")).
			Comment("接单"),
	}
}

func (User) Indexes() []ent.Index {
	return nil
	//return []ent.Index{
	//	index.Fields("id").
	//		Edges("receiver"),
	//}
	//return []ent.Index{
	//	index.Fields("phone").
	//		Unique(),
	//}
}
