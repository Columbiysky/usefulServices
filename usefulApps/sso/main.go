package main

import (
	"log"
	"sso/controllers/accountController"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	readDocs()
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/getAccount/", accountController.GetAccount)
	router.POST("/registerAccount/", accountController.RegisterAccount)
	router.POST("/registerActivity/", accountController.RegisterActivity)

	router.Run()
}

func readDocs() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
		return
	}
}
