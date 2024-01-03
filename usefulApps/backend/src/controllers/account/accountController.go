package controllers

import (
	accLogic "main/logic/account"
	models "main/models"

	"github.com/gin-gonic/gin"
)

func GetAccount(context *gin.Context) {
	RegisterAccount(context)
}

func RegisterAccount(context *gin.Context) {
	var acc models.Account

	// context.BindJSON(&acc)
	accLogic.RegisterAccount(acc)
}
