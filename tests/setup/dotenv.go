package tests

import "github.com/joho/godotenv"

func SetupDotEnvFile() {
	godotenv.Load("../.env.test")
}
