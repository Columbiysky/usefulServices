package dbLogic

import (
	"log"
	"sso/dbLogic/base"
	"sso/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Login(login string, pass string) *models.Account {
	dbConn := getConnection()

	var account models.Account
	dbConn.Where(map[string]interface{}{"login": login, "password": pass}).First(&account)

	if account.Id != 0 {
		return &account
	}

	log.Println("Account not found. Wrong login or password")
	return nil
}

func getConnection() *gorm.DB {
	dbConn, err := gorm.Open(postgres.Open(base.GetConnectionToSSOAccounts()), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return dbConn
}
