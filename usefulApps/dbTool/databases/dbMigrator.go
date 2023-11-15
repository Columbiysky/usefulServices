package databases

import (
	"dbTool/databases/db_internal"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Id       int64  `gorm:"not null" json:"account_id"`
	Login    string `gorm:"not null" json:"account_login"`
	Password string `gorm:"not null" json:"account_password"`
}

type Token struct {
	gorm.Model
	Id         int64  `gorm:"not null" json:"token_id"`
	TokenValue string `gorm:"not null" json:"token_value"`
}

type AccountToken struct {
	gorm.Model
	Id      int64   `gorm:"not null" json:"account_token_id"`
	Account Account `gorm:"not null" json:"account_token_account"`
	Token   Token   `gorm:"not null" json:"account_token_token"`
}

func Migrate() {
	db, err := gorm.Open(postgres.Open(db_internal.GetConnectionToSSOAccounts()), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return
	}

	err = db.AutoMigrate(&Account{}, &Token{}, &AccountToken{})

	if err != nil {
		log.Fatalln(err)
		return
	}

	db.Set("gorm:table_options", "ENGINE=PostgreSQL")

	log.Println("Migration completed sucsessfully")
}
