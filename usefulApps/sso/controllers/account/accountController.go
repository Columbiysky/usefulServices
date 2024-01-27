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

	if tokenValue != nil {
		context.JSON(http.StatusOK, tokenValue)
		return
	}

	errorMessage := models.ErrorMessage{Message: "Account not found. Wrong login or password"}
	context.JSON(http.StatusNotFound, errorMessage)
}
