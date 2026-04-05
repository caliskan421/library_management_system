package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammetalicaliskan/libranet/internal/dto"
	"github.com/muhammetalicaliskan/libranet/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// POST /api/auth/register
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Geçersiz istek verisi",
		})
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Ad, e-posta ve şifre alanları zorunludur",
		})
	}

	if len(req.Password) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Şifre en az 8 karakter olmalıdır",
		})
	}

	result, err := h.authService.Register(c.Context(), req)
	if err != nil {
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: "Kayıt işlemi sırasında bir hata oluştu",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

// POST /api/auth/login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Geçersiz istek verisi",
		})
	}

	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "E-posta ve şifre alanları zorunludur",
		})
	}

	result, err := h.authService.Login(c.Context(), req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Message: err.Error(),
			})
		}
		if errors.Is(err, service.ErrAccountLocked) {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: "Giriş işlemi sırasında bir hata oluştu",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
