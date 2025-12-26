package middleware

import "github.com/gofiber/fiber/v2"

func AdminOnly(c *fiber.Ctx) error {
	if c.Locals("role") != "admin" {
		return c.Status(403).SendString("Admin only")
	}
	return c.Next()
}
