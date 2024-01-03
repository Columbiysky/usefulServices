package controllers

import (
	"log"
	"net/http"

	accLogic "sso/logic/account"
	"sso/models"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var acc models.AccountLoginDbo

	if err := context.BindJSON(&acc); err != nil {
		log.Println(err)
	}

	tokenValue := accLogic.Login(acc)
	context.JSON(http.StatusOK, tokenValue)
}

func RegisterActivity(context *gin.Context) {
	log.Printf("RegisterActivity")
}
