package logic

import (
	"sso/models"
	"testing"

	accLogic "sso/logic/account"
	tokenLogic "sso/logic/token"
)

func TestLogin(t *testing.T) {
	testAcc := models.AccountLoginDbo{
		Login:    "t",
		Password: "t",
	}

	res := accLogic.Login(testAcc)
	if res == nil {
		t.Error("Result was incorrect, got nil")
	}

	cleanup()
}

func cleanup() {
	tokenLogic.CleanOldTokens()
}
