package dbLogic

import (
	"backend/dbLogic/base"
	"backend/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetAccountById(id int64) *models.Account {
	dbConn := getConnection()

	var account models.Account
	dbConn.First(&account, id)

	if account.Id != 0 {
		return &account
	}

	log.Println("Account not found. Wrong id")
	return nil
}

func RegisterAccount(account models.Account) int {
	dbConn := getConnection()
	dbConn.Create(&account)
	return account.Id
}

func GetAccountByLoginAndPassword(login string, pass string) *models.Account {
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
