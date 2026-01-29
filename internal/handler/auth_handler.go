package handler

import (
	"backend-boking-ticket/internal/entity"
	"backend-boking-ticket/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req entity.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entity.Response{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	response, err := h.service.Register(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entity.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(entity.Response{
		Status:  "sukses",
		Message: "Registrasi berhasil",
		Data:    response,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req entity.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entity.Response{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	response, err := h.service.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(entity.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(entity.Response{
		Status:  "sukses",
		Message: "Login berhasil",
		Data:    response,
	})
}
