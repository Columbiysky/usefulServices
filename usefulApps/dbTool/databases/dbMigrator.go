package databases

import (
	"dbTool/databases/db_internal"
	"dbTool/databases/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() {
	db, err := gorm.Open(postgres.Open(db_internal.GetConnectionToSSOAccounts()), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return
	}

	err = db.AutoMigrate(&models.Account{}, &models.Token{}, &models.AccountToken{})

	if err != nil {
		log.Fatalln(err)
		return
	}

	db.Set("gorm:table_options", "ENGINE=PostgreSQL")

	log.Println("Migration completed sucsessfully")
}
