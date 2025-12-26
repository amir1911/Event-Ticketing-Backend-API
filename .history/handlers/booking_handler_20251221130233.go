package handlers

import (


	"github.com/gofiber/fiber/v2"
)

type Booking struct {
	ID      int `json:"id"`
	EventID int `json:"event_id"`
	Qty     int `json:"qty"`
}

var bookings = []Booking{}

func CreateBooking(c *fiber.Ctx) error {
	req := new(Booking)
	c.BodyParser(req)

	for i, e := range events {
		if e.ID == req.EventID {

			// 🔒 GATEKEEPER LOGIC
			if req.Qty > e.Ticket {
				return c.Status(400).SendString("Ticket not enough")
			}

			events[i].Ticket -= req.Qty
			req.ID = len(bookings) + 1
			bookings = append(bookings, *req)

			return c.JSON(req)
		}
	}
	return c.Status(404).SendString("Event not found")
}

func GetBookings(c *fiber.Ctx) error {
	return c.JSON(bookings)
}
