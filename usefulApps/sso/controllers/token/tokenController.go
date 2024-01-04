package controllers

import (
	"log"
	"net/http"
	tokenLogic "sso/dbLogic/token"
	"strings"

	"github.com/gin-gonic/gin"
)

type TokenValueStruct struct {
	Value string `json:"value"`
}

func CheckToken(context *gin.Context) {
	bearerToken := context.Request.Header.Get("Authorization")
	bearerTokenSplitted := strings.Split(bearerToken, " ")

	if len(bearerTokenSplitted) < 2 {
		log.Println("Account unauthorized")
		context.JSON(http.StatusUnauthorized, gin.H{
			"status": "not exists",
		})
		return
	}

	tokenValue := strings.Split(bearerToken, " ")[1]
	result := tokenLogic.GetToken(tokenValue)
	if result == nil {
		log.Println("Account unauthorized")
		context.JSON(http.StatusUnauthorized, gin.H{
			"status": "not exists",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "exists"})
}

func RegisterActivity(context *gin.Context) {
	var tokenValue TokenValueStruct
	context.BindJSON(&tokenValue)

	tokenLogic.RegisterActivity(tokenValue.Value)
}
