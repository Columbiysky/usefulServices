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
	if token.Id == 0 {
		log.Println("Token doesn't exist")
		return
	}
	token.LastActivityTime = time.Now().UTC()
	dbConn.Save(&token)

}

func RegisterToken(accountId int, tokenValue string) {
	dbConn := getConnection()

	token := models.Token{
		Id:               0,
		TokenValue:       tokenValue,
		LastActivityTime: time.Now().UTC(),
	}

	dbConn.Create(&token)

	accountToken := models.AccountToken{
		Id:        0,
		AccountId: accountId,
		TokenId:   token.Id,
	}
	dbConn.Create(&accountToken)
}

func getConnection() *gorm.DB {
	dbConn, err := gorm.Open(postgres.Open(base.GetConnectionToSSOAccounts()), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return dbConn
}
