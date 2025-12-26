package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go-belajar/config"
	"go-belajar/models"
)

func CreateEvent(c *fiber.Ctx) error {
	var event models.Event
	c.BodyParser(&event)

	config.DB.Create(&event)
	return c.JSON(event)
}

func GetEvents(c *fiber.Ctx) error {
	var events []models.Event
	config.DB.Find(&events)
	return c.JSON(events)
}

func GetEventByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event

	if err := config.DB.First(&event, id).Error; err != nil {
		return c.Status(404).SendString("Event not found")
	}

	return c.JSON(event)
}

func UpdateEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event

	if err := config.DB.First(&event, id).Error; err != nil {
		return c.Status(404).SendString("Event not found")
	}

	c.BodyParser(&event)
	config.DB.Save(&event)

	return c.JSON(event)
}

func DeleteEvent(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	config.DB.Delete(&models.Event{}, id)
	return c.SendString("Event deleted")
}
