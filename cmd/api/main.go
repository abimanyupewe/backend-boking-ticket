package main

import (
	"backend-boking-ticket/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Setup Routes
	routes.SetupRoutes(app)

	// Start server on port 3000
	log.Fatal(app.Listen(":3000"))
}
