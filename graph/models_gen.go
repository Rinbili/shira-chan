// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

type SignInput struct {
	// 用户名
	Uname *string `json:"uname,omitempty"`
	// 密码
	Passwd string `json:"passwd"`
	// 手机号码
	Phone string `json:"phone"`
}

type Token struct {
	Token       *string `json:"token,omitempty"`
	IsAdmin     *bool   `json:"is_admin,omitempty"`
	IsSecretary *bool   `json:"is_secretary,omitempty"`
}
