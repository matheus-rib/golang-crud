package main

import (
	"log"

	"github.com/matheus-rib/golang-crud/config"
	"github.com/matheus-rib/golang-crud/database"
)

func main() {
	config.SetupDotEnvFile()
	database.DB = config.ConnectDatabase()
	router := config.SetupRouter()

	log.Fatal(router.Run(config.GetAPIAddress()))
}
