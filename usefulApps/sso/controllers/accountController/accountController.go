package accountController

import (
	"log"

	"sso/logic"
	"sso/models"

	"github.com/gin-gonic/gin"
)

func GetAccount(context *gin.Context) {
	RegisterAccount(context)
}

func RegisterAccount(context *gin.Context) {
	var acc models.Account

	// context.BindJSON(&acc)
	logic.RegisterAccount(acc)
}

func Login(context *gin.Context) {
	logic.Login("t", "t")
}

func RegisterActivity(context *gin.Context) {
	log.Printf("RegisterActivity")
}
