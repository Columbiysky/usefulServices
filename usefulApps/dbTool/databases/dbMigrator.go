package databases

import (
	"dbTool/databases/db_internal"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Id       int64  `gorm:"primaryKey;not null;autoIncrement:1;autoIncrementIncrement:1" json:"account_id"`
	Login    string `gorm:"not null" json:"account_login"`
	Password string `gorm:"not null" json:"account_password"`
}

type Token struct {
	gorm.Model
	Id         int64  `gorm:"primaryKey;not null;autoIncrement:1;autoIncrementIncrement:1" json:"token_id"`
	TokenValue string `gorm:"not null" json:"token_value"`
}

type AccountToken struct {
	gorm.Model
	Id           int64 `gorm:"primaryKey;not null;autoIncrement:1;autoIncrementIncrement:1" json:"account_token_id"`
	AccountRefer int64
	TokenRefer   int64
	Account      Account `gorm:"foreignKey:AccountRefer"`
	Token        Token   `gorm:"foreignKey:TokenRefer"`
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
