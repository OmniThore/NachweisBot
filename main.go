package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type User struct {
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_NAME string
}

func main() {
}

func ReadEnv() User {
	// Get credentials from the environment
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return User{
		DB_USER: os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASS"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_NAME: os.Getenv("DB_NAME"),
	}
}
