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

func GetToken(id int64) {

}

func RegisterActivity(token string) {

}
