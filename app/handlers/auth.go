package handlers

import (
	"github.com/gin-gonic/gin"
	"shira-chan-dev/app/utils"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//r := Response{}
		tokenString := c.Request.Header.Get("Authorization")
		if len(tokenString) != 0 {
			tokenString = tokenString[7:]
			claims, err := utils.ParseToken(tokenString)
			if err == nil {
				u, err := utils.Client.User.Get(c.Copy(), claims.UId)
				//不接受用户数据修改前签发的token
				if err == nil && u.IsActive == true && u.UpdatedAt.Before(claims.IssuedAt.Time) {
					//token有效可用
					c.Set("is_authed", true)
					c.Set("is_admin", u.IsAdmin)
					c.Next()
					return
				}
			}
		}
		c.Set("is_authed", false)
		c.Next()
		//r.Code = http.StatusUnauthorized
		//r.err = errors.New("unauthorized")
		//ResponseJSON(c, r)
		//c.Abort()
	}
}
