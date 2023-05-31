package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := Response{}
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString != "" {
			//claims, err := utils.ParseToken(tokenString)
			//if err == nil {
			//	c.Request.Header.Set("uid", strconv.Itoa(claims.UId))
			//	c.Request.Header.Set("is_admin", strconv.FormatBool(claims.IsAdmin))
			//	c.Next()
			//}
			c.Next()
		}
		r.Code = http.StatusUnauthorized
		r.err = errors.New("unauthorized")
		ResponseJSON(c, r)
		c.Abort()
	}
}
