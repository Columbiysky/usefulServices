package token

import (
	"log"
	"sso/dbLogic/base"
	"sso/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RegisterActivity(tokenValue string) {
	dbConn := getConnection()

	var token models.Token
	dbConn.Where("token_value = ?", tokenValue).First(&token)
	token.LastActivityTime = time.Now().UTC()
	dbConn.Save(&token)
}

func getConnection() *gorm.DB {
	dbConn, err := gorm.Open(postgres.Open(base.GetConnectionToSSOAccounts()), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return dbConn
}
