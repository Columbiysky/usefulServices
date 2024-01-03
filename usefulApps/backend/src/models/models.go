package models

type Account struct {
	Id       int    `gorm:"not null" json:"account_id"`
	Login    string `gorm:"not null" json:"account_login"`
	Password string `gorm:"not null" json:"account_password"`
}

func (Account) TableName() string {
	return "sso.accounts"
}
