package base

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func GetConnectionToSSOAccounts() string {
	viper.AddConfigPath(".")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=%s TimeZone=%s",
		viper.Get("db.host"),
		viper.Get("db.user"),
		viper.Get("db.password"),
		viper.Get("db.dbNames.sso_accounts"),
		viper.Get("db.port"),
		viper.Get("db.sslmode"),
		viper.Get("db.TimeZone"),
	)
	return dsn
}
