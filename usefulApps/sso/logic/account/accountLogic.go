package logic

import (
	accDb "sso/dbLogic/account"
	token "sso/logic/token"
	"sso/models"
)

func Login(acc models.AccountLoginDbo) *models.TokenValueDto {
	account := accDb.GetAccountByLoginAndPassword(acc.Login, acc.Password)

	if account != nil {
		res := token.GenerateTokenForAccount(account.Id, account.Login)
		token.RegisterToken(account.Id, res)
		return &models.TokenValueDto{
			TokenValue: res,
		}
	}

	return nil
}
