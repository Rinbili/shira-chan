package rule

import (
	"context"
	"shira-chan-dev/ent/privacy"
	"shira-chan-dev/ent/rule/viewer"
)

func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view, _ := viewer.FromContext(ctx)
		if view == nil {
			return privacy.Denyf("user-context is missing")
		}
		// Skip to the next privacy rule (equivalent to returning nil).
		return privacy.Skip
	})
}

// AllowIfAdmin is a rule that returns Allow decision if the viewer is admin.
func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view, err := viewer.FromContext(ctx)
		if err != nil {
			return privacy.Denyf("user-context is missing")
		}
		if viewer.IsAdmin(view) {
			return privacy.Allow
		}
		// Skip to the next privacy rule (equivalent to returning nil).
		return privacy.Skip
	})
}

// AllowIfSecretary is a rule that returns Allow decision if the viewer is admin.
func AllowIfSecretary() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view, err := viewer.FromContext(ctx)
		if err != nil {
			return privacy.Denyf("user-context is missing")
		}
		if viewer.IsSecretary(view) {
			return privacy.Allow
		}
		// Skip to the next privacy rule (equivalent to returning nil).
		return privacy.Skip
	})
}
