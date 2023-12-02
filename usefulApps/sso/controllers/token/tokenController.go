package controllers

import (
	tokenLogic "sso/dbLogic/token"

	"github.com/gin-gonic/gin"
)

type TokenValueStruct struct {
	Value string `json:"value"`
}

func RegisterActivity(context *gin.Context) {
	var tokenValue TokenValueStruct
	context.BindJSON(&tokenValue)

	tokenLogic.RegisterActivity(tokenValue.Value)
}
