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
		Id:       1,
		Login:    "testLogin",
		Password: "testPassword",
	}

	dbConn.Create(&dbEntity)

	return dbEntity.Id
}

func GetToken(id int64) {

}

func RegisterActivity(token string) {

}
