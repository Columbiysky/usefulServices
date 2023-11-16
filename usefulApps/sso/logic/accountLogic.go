package logic

import (
	"sso/dbLogic"
	"sso/models"
)

func GetAccount(id int64) {

}

func RegisterAccount(account models.Account) int64 {
	id := dbLogic.RegisterAccount(account)
	return id
}

func Login(login string, pass string) string {
	res := dbLogic.Login(login, pass)
	return res
}

func GetToken(id int64) {

}

func RegisterActivity(token string) {

}
