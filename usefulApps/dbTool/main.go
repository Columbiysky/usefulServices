package main

import (
	"dbTool/databases"
)

func main() {
	databases.CreateDatabases()
	databases.Migrate()
}
