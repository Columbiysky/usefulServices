package logic

import (
	tokenDb "backend/dbLogic/token"
	"backend/models"
)

func GetToken(tokenValue string) *models.Token {
	token := tokenDb.GetToken(tokenValue)
	return token
}
