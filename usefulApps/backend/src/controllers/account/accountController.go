package controllers

import (
	accLogic "backend/logic/account"
	models "backend/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAccountById(context *gin.Context) {
	if !checkToken(context) {
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

func checkToken(context *gin.Context) bool {
	bearerToken := context.Request.Header.Get("Authorization")
	bearerTokenSplitted := strings.Split(bearerToken, " ")

	if len(bearerTokenSplitted) < 2 {
		return false
	}

	tokenValue := strings.Split(bearerToken, " ")[1]
	req, err := http.NewRequest("GET", "http://127.0.0.1:8081/checkToken", nil)

	if err != nil {
		log.Println(err)
		return false
	}

	req.Header.Add("Authorization", "Bearer "+tokenValue)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var dat map[string]interface{}
	json.Unmarshal([]byte(body), &dat)
	status := dat["status"].(string)

	if status == "exists" {
		return true
	} else {
		return false
	}
}
