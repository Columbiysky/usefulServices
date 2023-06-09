package controllers

import (
	logic "main/logic/hello"
	"main/logic/hello/params"

	"github.com/gin-gonic/gin"
)

func Hello(context *gin.Context) {
	logic.Hello(context)
}

func HelloObject(context *gin.Context) {
	var obj params.HelloObject

	context.BindJSON(&obj)

	logic.HelloObject(context, obj)
}
