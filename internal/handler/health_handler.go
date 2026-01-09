package handler

import (
	"backend-boking-ticket/internal/entity"

	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(c *fiber.Ctx) error {
	return c.JSON(entity.Response{
		Status:  "sukses",
		Message: "Sistem berjalan normal (Clean Architecture)",
		Data:    nil,
	})
}
