package main

import (
	"log"

	_ "ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	log.Println("its a migrator")
	// databases.CreateDatabases()
	// databases.Migrate()
}
