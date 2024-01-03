package logic

import (
	accDb "main/dbLogic/account"
	"main/models"
)

func GetAccount(id int64) {

}

func RegisterAccount(account models.Account) int {
	id := accDb.RegisterAccount(account)
	return id
}
