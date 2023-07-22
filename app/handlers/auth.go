package handlers

import (
	"github.com/gin-gonic/gin"
	"shira-chan-dev/app/utils"
	"shira-chan-dev/ent/privacy"
)

// AuthHandler
// @Description: 中间件：验证token
// @return gin.HandlerFunc:
func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if len(tokenString) != 0 {
			tokenString = tokenString[7:]
			claims, err := utils.ParseToken(tokenString)
			if err == nil {
				u, err := utils.Client.User.Get(privacy.DecisionContext(c, privacy.Allow), claims.UId)
				//不接受用户数据修改前签发的token
				if err == nil && u.IsActive == true && u.UpdatedAt < claims.IssuedAt.Time.Unix() {
					//token有效可用
					c.Set("user", u)
					c.Next()
					return
				}
			}
		}
		c.Set("user", nil)
		c.Next()
	}
}
