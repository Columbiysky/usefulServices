package main

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello/", controllers.Hello)
	router.POST("/hello/", controllers.HelloObject)
	router.Run()
}
