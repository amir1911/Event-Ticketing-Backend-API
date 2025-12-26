package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"go-belajar/database"
	"go-belajar/models"
)

func Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).SendString("Invalid input")
	}

	var user models.User
	if err := database.DB.
		Where("email = ?", input.Email).
		First(&user).Error; err != nil {
		return c.Status(401).SendString("Invalid credentials")
	}

	if input.Password != user.Password {
		return c.Status(401).SendString("Invalid credentials")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).SendString("Failed to sign token")
	}

	return c.JSON(fiber.Map{
		"token": signed,
	})
}
