package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"shira-chan-dev/app/utils"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := Response{}
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString != "" {
			claims, err := utils.ParseToken(tokenString)
			if err == nil {
				u, err := utils.Client.User.Get(c.Copy(), claims.UId)
				//不接受用户数据修改前签发的token
				if err != nil && u.IsActive == true && u.UpdatedAt.Before(claims.IssuedAt.Time) {
					c.Next()
					return
				}
			}
		}
		r.Code = http.StatusUnauthorized
		r.err = errors.New("unauthorized")
		ResponseJSON(c, r)
		c.Abort()
	}
}
