package controllers

import (
	"log"

	accLogic "sso/logic/account"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	accLogic.Login("t", "t")
}

func RegisterActivity(context *gin.Context) {
	log.Printf("RegisterActivity")
}
