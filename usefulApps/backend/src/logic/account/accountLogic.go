package logic

import (
	accDb "backend/dbLogic/account"
	"backend/models"
	"log"
)

func GetAccountById(id int64) *models.Account {
	account := accDb.GetAccountById(id)
	return account
}

func RegisterAccount(account models.Account) int {
	if account.Login == "" || account.Password == "" {
		log.Println("Account login and password must be not empty")
		return -1
	}

	newId := accDb.RegisterAccount(account)
	return newId
}
