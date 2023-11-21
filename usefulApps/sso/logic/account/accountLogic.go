package logic

import (
	accDb "sso/dbLogic/account"
	token "sso/logic/token"
	"sso/models"
)

func GetAccount(id int64) {

}

func RegisterAccount(account models.Account) int {
	id := accDb.RegisterAccount(account)
	return id
}

func Login(login string, pass string) string {
	account := accDb.GetAccountByLoginAndPassword(login, pass)

	if account != nil {
		res := token.GenerateTokenForAccount(account.Id, account.Login)
		return res
	}

	return ""
}

func GetToken(id int64) {

}

func RegisterActivity(token string) {

}
