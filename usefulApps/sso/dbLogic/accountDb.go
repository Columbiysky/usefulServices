package dbLogic

import (
	"log"
	"sso/dbLogic/base"
	"sso/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetAccount(id int64) {

}

func RegisterAccount(account models.Account) int64 {
	dbConn, err := gorm.Open(postgres.Open(base.GetConnectionToSSOAccounts()), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return -1
	}

	dbEntity := models.Account{
		Login:    "t",
		Password: "t",
	}

	dbConn.Create(&dbEntity)

	return dbEntity.Id
}

func Login(login string, pass string) string {
	dbConn, err := gorm.Open(postgres.Open(base.GetConnectionToSSOAccounts()), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return ""
	}

	var account models.Account
	dbConn.Where(map[string]interface{}{"login": login, "password": pass}).First(&account)

	if account.Id != 0 {
		return "t"
	}

	return "f"
}

func GetToken(id int64) {

}

func RegisterActivity(token string) {

}
