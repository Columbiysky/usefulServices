package controllers

import (
	tokenLogic "sso/logic/token"

	"github.com/gin-gonic/gin"
)

func RegisterActivity(context *gin.Context) {
	var tokenValue string

	// context.BindJSON(&tokenValue)
	tokenLogic.RegisterActivity(tokenValue)
}
