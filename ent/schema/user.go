package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"golang.org/x/crypto/bcrypt"
	gen "shira-chan-dev/ent"
	"shira-chan-dev/ent/hook"
	"shira-chan-dev/ent/privacy"
	"shira-chan-dev/ent/rule"
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
			Through("receives", Receive.Type).
			Comment("接单"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			// 增、改：密码加密
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
					// 获取密码字段的值
					passwd, ok := m.Passwd()
					if !ok {
						// 如果字段不存在或类型不正确，直接调用下一个处理器
						return next.Mutate(ctx, m)
					}
					temp, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
					if err != nil {
						return nil, err
					} else {
						m.SetPasswd(string(temp))
					}
					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}

func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoViewer(),
			rule.AllowIfAdmin(),
			rule.UserSelfCheckRule(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			rule.DenyIfNoViewer(),
			rule.AllowIfSecretary(),
			rule.UserSelfCheckRule(),
			privacy.AlwaysDenyRule(),
		},
	}
}
