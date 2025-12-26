package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-belajar/config"
	"go-belajar/models"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	c.BodyParser(&user)

	config.DB.Create(&user)
	return c.JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	config.DB.Find(&users)
	return c.JSON(users)
}
