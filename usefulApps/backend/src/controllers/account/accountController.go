package controllers

import (
	accLogic "backend/logic/account"
	tokenLogic "backend/logic/token"
	models "backend/models"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAccountById(context *gin.Context) {
	if checkToken(context) == nil {
		log.Println("Account unauthorized")
		context.JSON(http.StatusUnauthorized, gin.H{
			"reason": "unauthorized",
		})
		return
	}

	parsedQuery := context.Request.URL.Query()
	id, err := strconv.ParseInt(parsedQuery["id"][0], 10, 10)

	if err != nil {
		log.Println(err)
	}

	result := accLogic.GetAccountById(id)
	context.JSON(http.StatusOK, result)
}

func RegisterAccount(context *gin.Context) {
	var acc models.Account

	if err := context.BindJSON(&acc); err != nil {
		log.Println(err)
	}

	result := accLogic.RegisterAccount(acc)
	context.JSON(http.StatusOK, result)
}

func checkToken(context *gin.Context) *models.Token {
	bearerToken := context.Request.Header.Get("Authorization")
	bearerTokenSplitted := strings.Split(bearerToken, " ")

	if len(bearerTokenSplitted) < 2 {
		return nil
	}

	tokenValue := strings.Split(bearerToken, " ")[1]
	result := tokenLogic.GetToken(tokenValue)
	return result
}
