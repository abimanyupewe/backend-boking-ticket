package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() Config {
	// Load .env file, ignore error if file not found (might be production)
	err := godotenv.Load()
	if err != nil {
		log.Println("Note: .env file not found, using environment variables")
	}

	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
