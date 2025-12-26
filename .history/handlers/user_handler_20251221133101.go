package handlers

import (
	"strconv"

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

func GetUserByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).SendString("User not found")
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).SendString("User not found")
	}

	c.BodyParser(&user)
	config.DB.Save(&user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	config.DB.Delete(&models.User{}, id)
	return c.SendString("User deleted")
}
