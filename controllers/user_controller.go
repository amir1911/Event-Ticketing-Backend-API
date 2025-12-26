package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go-belajar/database"
	"go-belajar/models"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	database.DB.Create(&user)
	return c.JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid user ID")
	}

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).SendString("User not found")
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid user ID")
	}

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).SendString("User not found")
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	database.DB.Save(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid user ID")
	}

	database.DB.Delete(&models.User{}, id)
	return c.SendString("User deleted")
}
