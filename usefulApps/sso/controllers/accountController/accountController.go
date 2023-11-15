package accountController

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetAccount(context *gin.Context) {
	log.Printf("GetAccount")
}

func RegisterAccount(context *gin.Context) {
	log.Printf("RegisterAccount")
}

func RegisterActivity(context *gin.Context) {
	log.Printf("RegisterActivity")
}
