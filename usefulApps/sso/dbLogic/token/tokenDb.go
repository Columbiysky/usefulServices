package token

import (
	"log"
	"sso/dbLogic/base"
	"sso/models"
	"time"

	"github.com/samber/lo"
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

func CleanOldTokens() {
	dbConn := getConnection()

	timeForRequest := time.Now().Add(-24 * time.Hour)
	var tokens []models.Token
	dbConn.Where("last_activity_time < ?", timeForRequest).Find(&tokens)

	tokensIds := lo.Map(tokens, func(x models.Token, index int) int {
		return x.Id
	})
	var account_tokens models.AccountToken
	dbConn.Where("token_id in ?", tokensIds).Delete(&account_tokens)

	dbConn.Where("id in ?", tokensIds).Delete(&tokens)
}

func getConnection() *gorm.DB {
	dbConn, err := gorm.Open(postgres.Open(base.GetConnectionToSSOAccounts()), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return dbConn
}
