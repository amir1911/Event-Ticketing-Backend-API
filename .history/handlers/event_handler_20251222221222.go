package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go-belajar/database"
	"go-belajar/models"
)

func CreateEvent(c *fiber.Ctx) error {
	var event models.Event

	if err := c.BodyParser(&event); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	database.DB.Create(&event)
	return c.JSON(event)
}

func GetEvents(c *fiber.Ctx) error {
	var events []models.Event
	database.DB.Find(&events)
	return c.JSON(events)
}

func GetEventByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid event ID")
	}

	var event models.Event
	if err := database.DB.First(&event, id).Error; err != nil {
		return c.Status(404).SendString("Event not found")
	}

	return c.JSON(event)
}

func UpdateEvent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid event ID")
	}

	var event models.Event
	if err := database.DB.First(&event, id).Error; err != nil {
		return c.Status(404).SendString("Event not found")
	}

	if err := c.BodyParser(&event); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	database.DB.Save(&event)
	return c.JSON(event)
}

func DeleteEvent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid event ID")
	}

	database.DB.Delete(&models.Event{}, id)
	return c.SendString("Event deleted")
}
