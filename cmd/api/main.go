package main

import (
	"backend-boking-ticket/config"
	"backend-boking-ticket/internal/entity"
	"backend-boking-ticket/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 1. Load Config
	cfg := config.LoadConfig()

	// 2. Connect to Database
	db := config.ConnectDB(cfg.DatabaseURL)

	// Auto Migrate (Membuat tabel otomatis)
	db.AutoMigrate(&entity.Booking{}, &entity.User{})

	// Initialize Fiber app
	app := fiber.New()

	// Setup Routes
	routes.SetupRoutes(app, db)

	// Start server on port 3000
	log.Fatal(app.Listen(":3000"))
}
