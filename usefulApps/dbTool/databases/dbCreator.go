package databases

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

func CreateDatabases() {
	readConfig()
	createDatabaseSSO_ACCOUNTS()
}

func createDatabaseSSO_ACCOUNTS() {
	dbName := "sso_accounts"
	db, err := gorm.Open(postgres.Open(buildConnectionString()), &gorm.Config{})

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

func readConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
		return
	}
}

func buildConnectionString() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s TimeZone=%s",
		viper.Get("db.host"),
		viper.Get("db.user"),
		viper.Get("db.password"),
		viper.Get("db.port"),
		viper.Get("db.sslmode"),
		viper.Get("db.TimeZone"),
	)
	return dsn
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
}
