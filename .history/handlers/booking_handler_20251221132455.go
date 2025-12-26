package handlers

import (
	"go-belajar/config"
	"go-belajar/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateBooking(c *fiber.Ctx) error {
	var req models.Booking
	c.BodyParser(&req)

	return config.DB.Transaction(func(tx *gorm.DB) error {
		var event models.Event

		if err := tx.First(&event, req.EventID).Error; err != nil {
			return c.Status(404).SendString("Event not found")
		}

		// 🔒 GATEKEEPER LOGIC
		if req.Qty > event.Ticket {
			return c.Status(400).SendString("Ticket not enough")
		}

		event.Ticket -= req.Qty
		tx.Save(&event)
		tx.Create(&req)

		return c.JSON(req)
	})
}
