package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/matheus-rib/golang-crud/config"
	"github.com/matheus-rib/golang-crud/database"
)

func main() {
	config.SetupDotEnvFile()
	database.DB = config.ConnectDatabase()

	router := gin.Default()

	config.RegisterRoutes(router)

	log.Fatal(router.Run(config.GetAPIAddress()))
}
