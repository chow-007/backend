package controllers

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func returnMsg(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(200, Response{
		Code:code,
		Data:data,
		Msg:msg,
	})
}

