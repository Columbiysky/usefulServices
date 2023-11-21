package token

import (
	"log"
	"sso/dbLogic/base"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getConnection() *gorm.DB {
	dbConn, err := gorm.Open(postgres.Open(base.GetConnectionToSSOAccounts()), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return dbConn
}
