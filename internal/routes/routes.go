package routes

import (
	"backend-boking-ticket/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	healthHandler := handler.NewHealthHandler()

	// Grouping routes bisa dilakukan di sini, misal /api/v1
	api := app.Group("/api")

	api.Get("/health", healthHandler.Check)

	// Route root tetap ada untuk tes sederhana
	app.Get("/", func(c *fiber.Ctx) error {
		return healthHandler.Check(c)
	})
}
