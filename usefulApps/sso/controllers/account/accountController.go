package controllers

import (
	"log"

	accLogic "sso/logic/account"
	"sso/models"

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

func Login(context *gin.Context) {
	accLogic.Login("t", "t")
}

func RegisterActivity(context *gin.Context) {
	log.Printf("RegisterActivity")
}
