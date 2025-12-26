package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ENVLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
