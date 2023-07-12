package graph

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
)

func HasPermission() func(ctx context.Context, obj interface{}, next graphql.Resolver, role Role) (interface{}, error) {
	return func(
		ctx context.Context,
		obj interface{},
		next graphql.Resolver,
		role Role,
	) (res interface{}, err error) {
		// 获取ginContext
		ginContext := ctx.Value("GinContextKey")
		if ginContext == nil {
			err := fmt.Errorf("could not retrieve gin.Context")
			return nil, err
		}
		gc, ok := ginContext.(*gin.Context)
		if !ok {
			err := fmt.Errorf("gin.Context has wrong type")
			return nil, err
		}
		// 是否登录
		if !gc.Value("is_authed").(bool) {
			return nil, errors.New("unauthorized")
		}
		switch role {
		case RoleAdmin:
			if !gc.Value("is_admin").(bool) {
				return nil, errors.New("permission denied")
			}
			break
		// 待完成！
		case RoleSecretary:
			break
		case RoleUser:
			break
		}
		// you can do your thing here for permissions
		return next(ctx)
	}
}
