package main

import (
	"main/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
		AllowHeaders:    []string{"Origin"},
	}))

	router.GET("/hello/", controllers.Hello)
	router.POST("/hello/", controllers.HelloObject)
	router.Run()
}
