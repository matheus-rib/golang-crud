package config

import (
	"fmt"
	"log"
	"os"

	"github.com/matheus-rib/golang-crud/domains/tasks"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDatabaseConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"))
}

func ConnectDatabase() *gorm.DB {
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  getDatabaseConnectionString(),
		PreferSimpleProtocol: true,
	}))

	if err != nil {
		log.Fatal("Failed to connect to database.")
	}

	err = database.AutoMigrate(&tasks.Task{})

	if err != nil {
		log.Fatal("Failed to run auto migrate.")
	}

	fmt.Println("Database connected successfully")
	return database
}
