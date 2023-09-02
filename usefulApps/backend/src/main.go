package main

import (
	"main/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// router.Use(cors.New(cors.Config{
	// 	AllowAllOrigins:  true,
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "Cache-Control", "X-Requested-With"},
	// 	AllowCredentials: true,
	// }))

	router.Use(cors.Default())

	router.GET("/hello/", controllers.Hello)
	router.POST("/hello/", controllers.HelloObject)

	router.GET("/helloFromExecutor/", controllers.HelloFromExecutor)

	router.Run()
}
