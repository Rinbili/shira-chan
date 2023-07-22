package rule

import (
	"context"
	"entgo.io/ent/entql"
	"shira-chan-dev/ent"
	"shira-chan-dev/ent/privacy"
	"shira-chan-dev/ent/rule/viewer"
	"shira-chan-dev/ent/user"
)

// UserSelfCheckRule 是自查权限规则函数
func UserSelfCheckRule() privacy.QueryMutationRule {
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		// 从上下文中获取当前用户的信息
		view, err := viewer.FromContext(ctx)
		if err != nil {
			return privacy.Skip
		}
		if userFilter, ok := f.(*ent.UserFilter); ok {
			userFilter.WhereID(entql.IntEQ(view.ID))
			return privacy.Allow
		} else {
			return privacy.Skip
		}
	})
}

// OrderSelfCheckRule 是自查权限规则函数
func OrderSelfCheckRule() privacy.QueryMutationRule {
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		// 从上下文中获取当前用户的信息
		view, err := viewer.FromContext(ctx)
		if err != nil {
			return privacy.Skip
		}
		if userFilter, ok := f.(*ent.OrderFilter); ok {
			userFilter.WhereHasRequesterWith(user.IDEQ(view.ID))
			return privacy.Allow
		} else {
			return privacy.Skip
		}
	})
}
