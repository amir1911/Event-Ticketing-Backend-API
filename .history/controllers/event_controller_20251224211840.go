package controllers

import (
	"go-belajar/config"
	"go-belajar/models"

	"github.com/gofiber/fiber/v2"
)

func CreateEvent(c *fiber.Ctx) error {
	var event models.Event
	c.BodyParser(&event)

	config.DB.Create(&event)
	return c.JSON(event)
}
