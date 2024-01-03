package logic

import (
	accDb "sso/dbLogic/account"
	token "sso/logic/token"
)

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
