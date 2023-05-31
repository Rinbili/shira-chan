package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"shira-chan-dev/app/utils"
	"shira-chan-dev/ent"
	"shira-chan-dev/ent/user"
)

func SignHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := Response{Code: http.StatusBadRequest}
		type SignData struct {
			Phone  *string `json:"phone"`
			Passwd *string `json:"passwd"`
			UName  *string `json:"uname"`
		}
		data := &SignData{}
		r.err = c.BindJSON(data)
		if r.err == nil && data.Phone != nil && data.Passwd != nil {
			var u *ent.User
			//查找用户
			u, r.err = utils.Client.User.Query().
				Where(user.PhoneEQ(*data.Phone)).
				Only(c.Copy())
			if r.err != nil {
				r.err = nil
				//试图注册
				if data.UName != nil {
					u, r.err = utils.Client.User.Create().
						SetUname(*data.UName).
						SetPhone(*data.Phone).
						SetPasswd(*data.Passwd).
						Save(c.Copy())
				} else {
					r.err = errors.New("bad request")
				}
			}
			if r.err == nil {
				var token string
				token, r.err = utils.GetToken(u.ID, u.IsAdmin)
				r.Data = gin.H{"token": token}
			}
		} else {
			r.err = errors.New("bad request")
		}
		if r.err == nil {
			r.Code = http.StatusOK
			r.Msg = "success"
		}
		ResponseJSON(c, r)
	}
}