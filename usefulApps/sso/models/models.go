package models

type Account struct {
	Id       int64  `gorm:"not null" json:"account_id"`
	Login    string `gorm:"not null" json:"account_login"`
	Password string `gorm:"not null" json:"account_password"`
}

type Token struct {
	Id         int64  `gorm:"not null" json:"token_id"`
	TokenValue string `gorm:"not null" json:"token_value"`
}

type AccountToken struct {
	Id             int64 `gorm:"not null" json:"account_token_id"`
	AccountReferId int64
	TokenReferId   int64
}
