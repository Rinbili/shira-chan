package graph

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
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
	return gc, nil
}

func IsAuthed(ctx context.Context) bool {
	gc, err := GinContextFromContext(ctx)
	if err == nil {
		return gc.Value("is_authed").(bool)
	}
	return false
}

func IsAdmin(ctx context.Context) bool {
	gc, err := GinContextFromContext(ctx)
	if err == nil {
		if gc.Value("is_authed").(bool) {
			return gc.Value("is_admin").(bool)
		}
	}
	return false
}
