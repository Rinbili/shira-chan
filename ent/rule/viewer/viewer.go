package viewer

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"shira-chan-dev/ent"
)

func IsAdmin(u *ent.User) bool {
	return u.IsAdmin
}

func IsSecretary(u *ent.User) bool {
	return u.IsSecretary || u.IsAdmin
}

// FromContext returns the Viewer stored in a context.
func FromContext(ctx context.Context) (*ent.User, error) {
	// 获取ginContext
	gc, ok := ctx.Value("GinContextKey").(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	} else {
		user := gc.Value("user")
		if user != nil {
			return user.(*ent.User), nil
		} else {
			return nil, nil
		}
	}
}
