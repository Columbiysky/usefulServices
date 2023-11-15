package databases

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

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

func CreateDatabases() {
	createDatabaseSSO_ACCOUNTS()
}

func createDatabaseSSO_ACCOUNTS() {
	readConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s TimeZone=%s",
		viper.Get("db.host"),
		viper.Get("db.user"),
		viper.Get("db.password"),
		viper.Get("db.port"),
		viper.Get("db.sslmode"),
		viper.Get("db.TimeZone"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return
	}

	var checkResult int8
	checkDbExists := "SELECT count(datname) FROM pg_catalog.pg_database where datname ='sso_accounts'"
	db.Raw(checkDbExists).Scan(&checkResult)

	if checkResult == 0 {
		createDatabaseCommand := "create database sso_accounts"
		db.Exec(createDatabaseCommand)
		return
	}

	log.Println("Database sso_accounts exists")
}
