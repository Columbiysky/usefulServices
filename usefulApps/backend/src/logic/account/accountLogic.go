package logic

import (
	accDb "backend/dbLogic/account"
	"backend/models"
)

func GetAccountById(id int64) *models.Account {
	account := accDb.GetAccountById(id)
	return account
}

func RegisterAccount(account models.Account) int {
	newId := accDb.RegisterAccount(account)
	return newId
}
