package models

import "time"

type Account struct {
	Id       int    `gorm:"not null" json:"account_id"`
	Login    string `gorm:"not null" json:"account_login"`
	Password string `gorm:"not null" json:"account_password"`
}

func (Account) TableName() string {
	return "sso.accounts"
}

type AccountLoginDbo struct {
	Login    string `gorm:"not null" json:"account_login"`
	Password string `gorm:"not null" json:"account_password"`
}

func (AccountLoginDbo) TableName() string {
	return "sso.accounts"
}

type Token struct {
	Id               int       `gorm:"not null" json:"token_id"`
	TokenValue       string    `gorm:"not null" json:"token_value"`
	LastActivityTime time.Time `gorm:"not null" json:"last_activity_time"`
}

func (Token) TableName() string {
	return "sso.tokens"
}

type TokenValueDto struct {
	TokenValue string `gorm:"not null" json:"token_value"`
}

type AccountToken struct {
	Id             int `gorm:"not null" json:"account_token_id"`
	AccountReferId int
	TokenReferId   int
}

func (AccountToken) TableName() string {
	return "sso.accounts_tokens"
}
