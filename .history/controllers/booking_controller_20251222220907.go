package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-belajar/database"
	"go-belajar/models"
	"gorm.io/gorm"
)

func CreateBooking(c *fiber.Ctx) error {
	var req models.Booking

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	return database.DB.Transaction(func(tx *gorm.DB) error {
		var event models.Event

		if err := tx.First(&event, req.EventID).Error; err != nil {
			return c.Status(404).SendString("Event not found")
		}

		if req.Qty > event.Ticket {
			return c.Status(400).SendString("Ticket not enough")
		}

		event.Ticket -= req.Qty

		if err := tx.Save(&event).Error; err != nil {
			return err
		}

		if err := tx.Create(&req).Error; err != nil {
			return err
		}

		return c.JSON(req)
	})
}

func GetBookings(c *fiber.Ctx) error {
	var bookings []models.Booking
	database.DB.Find(&bookings)
	return c.JSON(bookings)
}
