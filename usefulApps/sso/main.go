package main

import (
	accountController "sso/controllers/account"
	tokenController "sso/controllers/token"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/getAccount/", accountController.GetAccount)
	router.POST("/registerAccount/", accountController.RegisterAccount)
	router.GET("/login/", accountController.Login)
	router.POST("/registerActivity/", tokenController.RegisterActivity)

	router.Run()
}
