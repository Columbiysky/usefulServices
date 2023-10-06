package main

import (
	"main/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/hello/", controllers.Hello)
	router.POST("/hello/", controllers.HelloObject)

	router.GET("/helloFromPyServ/", controllers.HelloFromPyServ)

	router.POST("/upload", controllers.HelloFile)

	router.Run()
}
