package logic

import (
	accDb "main/dbLogic/account"
	"main/models"
)

func GetAccountById(id int64) *models.Account {
	account := accDb.GetAccountById(id)
	return account
}

func RegisterAccount(account models.Account) int {
	id := accDb.RegisterAccount(account)
	return id
}
