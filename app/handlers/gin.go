package handlers

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"status"`
	Msg  string `json:"msg"`
	err  error
	Data interface{} `json:"data"`
}

func ResponseJSON(c *gin.Context, response Response) {
	if response.Msg == "" && response.err != nil {
		response.Msg = response.err.Error()
	}
	c.JSON(response.Code, response)
}
