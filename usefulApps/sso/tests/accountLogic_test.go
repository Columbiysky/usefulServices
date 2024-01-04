package logic

import (
	"sso/models"
	"testing"

	accLogic "sso/logic/account"
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
}
