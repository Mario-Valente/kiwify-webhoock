package middleware

import (
	"strings"

	"github.com/Mario-Valente/kiwify-webhoock/internal/config"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddelware(c *fiber.Ctx) error {

	config := config.NewConfig()

	if config.AuthSecret == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "auth secret is not set",
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing token authorization",
		})
	}

	parts := strings.Split(token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token format",
		})
	}
	if parts[1] != config.AuthSecret {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	return c.Next()
}
