package main

import (
	"log"
	accountController "sso/controllers/account"
	tokenController "sso/controllers/token"
	tokenLogic "sso/logic/token"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron/v2"
)

func main() {
	s, err := gocron.NewScheduler()

	if err != nil {
		log.Fatalln(err)
	}

	s.NewJob(
		gocron.DurationJob(
			24*time.Hour,
		),
		gocron.NewTask(cleanOldTokens),
	)

	s.Start()

	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/login/", accountController.Login)
	router.POST("/registerActivity/", tokenController.RegisterActivity)
	router.GET("/checkToken", tokenController.CheckToken)

	router.Run("127.0.0.1:8081")
}

func cleanOldTokens() {
	tokenLogic.CleanOldTokens()
}
