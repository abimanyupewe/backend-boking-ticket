package handler

import (
	"backend-boking-ticket/internal/entity"
	"backend-boking-ticket/internal/service"

	"github.com/gofiber/fiber/v2"
)

type BookingHandler struct {
	service service.BookingService
}

func NewBookingHandler(service service.BookingService) *BookingHandler {
	return &BookingHandler{
		service: service,
	}
}

func (h *BookingHandler) CreateBooking(c *fiber.Ctx) error {
	var req entity.BookingRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entity.Response{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	booking, err := h.service.CreateBooking(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(entity.Response{
		Status:  "sukses",
		Message: "Booking berhasil dibuat",
		Data:    booking,
	})
}

func (h *BookingHandler) GetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.service.GetBookingByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(entity.Response{
			Status:  "error",
			Message: "Booking tidak ditemukan",
		})
	}

	return c.JSON(entity.Response{
		Status:  "sukses",
		Message: "Data booking ditemukan",
		Data:    booking,
	})
}

func (h *BookingHandler) GetAllBookings(c *fiber.Ctx) error {
	bookings, err := h.service.GetAllBookings()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(entity.Response{
		Status:  "sukses",
		Message: "Data booking ditemukan",
		Data:    bookings,
	})
}
