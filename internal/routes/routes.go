package routes

import (
	"backend-boking-ticket/internal/handler"
	"backend-boking-ticket/internal/repository"
	"backend-boking-ticket/internal/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// ==============================
	// Dependency Injection / Wiring
	// ==============================

	// 1. Repository
	bookingRepo := repository.NewBookingRepository(db)
	userRepo := repository.NewUserRepository(db)

	// 2. Service
	bookingService := service.NewBookingService(bookingRepo)
	authService := service.NewAuthService(userRepo)

	// 3. Handler
	bookingHandler := handler.NewBookingHandler(bookingService)
	authHandler := handler.NewAuthHandler(authService)

	// ==============================
	// Routing Configuration
	// ==============================

	// Grouping routes /api
	api := app.Group("/api")

	// Auth Routes
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Booking Routes
	bookings := api.Group("/bookings")
	bookings.Post("/", bookingHandler.CreateBooking)
	bookings.Get("/", bookingHandler.GetAllBookings)
	bookings.Get("/:id", bookingHandler.GetBooking)
}
