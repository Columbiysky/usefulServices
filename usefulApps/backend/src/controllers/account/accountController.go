package controllers

import (
	"log"
	accLogic "main/logic/account"
	models "main/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAccountById(context *gin.Context) {
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
