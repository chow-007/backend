package controllers

import "github.com/gin-gonic/gin"

type Meta struct {
	Msg    string `json:"msg"`
	Status int   `json:"status"`
}
type Response struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

func returnMsg(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(200, Response{
		Data: data,
		Meta: Meta{Msg:msg, Status:code},
	})
}
