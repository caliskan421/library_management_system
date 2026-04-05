package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammetalicaliskan/libranet/internal/dto"
	jwtpkg "github.com/muhammetalicaliskan/libranet/pkg/jwt"
)

func Authenticate(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Message: "Authorization header is required",
			})
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Message: "Invalid authorization format. Use: Bearer <token>",
			})
		}

		claims, err := jwtpkg.ValidateToken(parts[1], jwtSecret)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Message: "Invalid or expired token",
			})
		}

		c.Locals("userId", claims.UserID)
		c.Locals("userEmail", claims.Email)
		c.Locals("userRole", claims.Role)

		return c.Next()
	}
}

func Authorize(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole, ok := c.Locals("userRole").(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Message: "Authentication required",
			})
		}

		for _, role := range roles {
			if userRole == role {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
			Message: "Bu işlem için yetkiniz bulunmamaktadır",
		})
	}
}
