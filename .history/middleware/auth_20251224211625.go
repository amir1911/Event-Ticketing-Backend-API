package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if auth == "" {
		return c.Status(401).SendString("Unauthorized")
	}

	token, err := jwt.Parse(auth[7:], func(t *jwt.Token) (interface{}, error) {
		return []byte("SECRET_KEY"), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).SendString("Invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Locals("user_id", claims["user_id"])
	c.Locals("role", claims["role"])

	return c.Next()
}
