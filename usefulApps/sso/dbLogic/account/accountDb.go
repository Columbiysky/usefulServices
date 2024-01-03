package dbLogic

import (
	"log"
	"sso/dbLogic/base"
	"sso/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetAccount(id int) {

}

func RegisterAccount(account models.Account) int {
	dbConn := getConnection()

	dbEntity := models.Account{
		Login:    "t",
		Password: "t",
	}

	dbConn.Create(&dbEntity)

	return dbEntity.Id
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

func GetToken(accountId int) string {
	dbConn := getConnection()
	var accountToken models.AccountToken
	var token models.Token
	dbConn.Where("account_id = ?", accountId).First(&accountToken)
	dbConn.Where("id = ?", accountToken.TokenReferId).First(&token)

	return token.TokenValue
}

func RegisterActivity(token string) {

}

func getConnection() *gorm.DB {
	dbConn, err := gorm.Open(postgres.Open(base.GetConnectionToSSOAccounts()), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return dbConn
}
