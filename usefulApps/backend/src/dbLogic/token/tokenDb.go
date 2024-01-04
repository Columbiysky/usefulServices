package dbLogic

import (
	"backend/dbLogic/base"
	"backend/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetToken(tokenValue string) *models.Token {
	dbConn := getConnection()

	var token models.Token
	dbConn.First(&token, "token_value = ?", tokenValue)

	if token.Id == 0 {
		log.Println("Token doesn't exist")
		return nil
	}

	var account_token models.AccountToken
	dbConn.First(&account_token, "token_id = ?", token.Id)

	if account_token.Id == 0 {
		log.Println("Not found an account which has this token")
		return nil
	}

	return &token
}

func getConnection() *gorm.DB {
	dbConn, err := gorm.Open(postgres.Open(base.GetConnectionToSSOAccounts()), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return dbConn
}
