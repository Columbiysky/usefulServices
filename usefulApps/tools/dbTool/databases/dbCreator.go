package databases

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"dbTool/databases/db_internal"
)

func CreateDatabases() {
	createDatabaseSSO_ACCOUNTS()
}

func createDatabaseSSO_ACCOUNTS() {
	dbName := "sso_accounts"
	db, err := gorm.Open(postgres.Open(db_internal.GetConnectionToDb()), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return
	}

	if !checkDbExists(db, dbName) {
		createDb(db, dbName)
		return
	}

	log.Printf("Database %s exists\n", dbName)
}

func checkDbExists(dbConnection *gorm.DB, dbName string) bool {
	var checkResult int8
	checkDbExists := fmt.Sprintf(
		"SELECT count(datname) FROM pg_catalog.pg_database where datname ='%s'",
		dbName,
	)
	dbConnection.Raw(checkDbExists).Scan(&checkResult)

	return checkResult > 0
}

func createDb(dbConnection *gorm.DB, dbName string) {
	createDatabaseCommand := fmt.Sprintf("create database %s", dbName)
	dbConnection.Exec(createDatabaseCommand)
	log.Printf("database %s has been created", dbName)
}
