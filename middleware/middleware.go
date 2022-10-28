package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token != "secret" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "Not have access or unauthenticated",
		})
	}
	return c.Next()
}