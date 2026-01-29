package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(url string) *gorm.DB {
	if url == "" {
		log.Fatal("DATABASE_URL must be set")
	}

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to database successfully")

	return db
}
