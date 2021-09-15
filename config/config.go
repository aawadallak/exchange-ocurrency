package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
