package logic

import (
	"main/logic/hello/params"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(context *gin.Context) {
	context.JSON(http.StatusOK, "hello")
	// test pull request
}

func HelloObject(context *gin.Context, obj params.HelloObject) {
	context.JSON(http.StatusOK, obj)
}
