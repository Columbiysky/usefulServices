package main

import (
	"sso/controllers/accountController"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/getAccount/", accountController.GetAccount)
	router.POST("/registerAccount/", accountController.RegisterAccount)
	router.POST("/registerActivity/", accountController.RegisterActivity)
	router.GET("/login/", accountController.Login)

	router.Run()
}
