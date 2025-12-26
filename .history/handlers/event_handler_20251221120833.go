package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Event struct {
	ID       int
	Title    string
	Date     string
	Location string
	Ticket   int
}

// Simulasi database (sementara)
var events = []Event{
	{ID: 1, Title: "Konser Musik Indie", Date: "20 Agustus 2025", Location: "Palembang", Ticket: 120},
	{ID: 2, Title: "Tech Conference", Date: "10 September 2025", Location: "Jakarta", Ticket: 50},
}

func EventPage(c *fiber.Ctx) error {
	return c.Render("events", fiber.Map{
		"Title":  "Daftar Event",
		"Events": events,
	}, "layout")
}

// DETAIL EVENT
func EventDetail(c *fiber.Ctx) error {
	return c.SendString("DETAIL EVENT MASUK")
}

// BOOKING + GATEKEEPER LOGIC
func BookTicket(c *fiber.Ctx) error {
	eventID, _ := strconv.Atoi(c.FormValue("event_id"))
	qty, _ := strconv.Atoi(c.FormValue("qty"))

	for i, e := range events {
		if e.ID == eventID {

			// GATEKEEPER LOGIC
			if qty > e.Ticket {
				return c.SendString("❌ Tiket tidak cukup (Sold Out)")
			}

			// Kurangi stok
			events[i].Ticket -= qty

			return c.SendString("✅ Booking berhasil!")
		}
	}

	return c.SendString("Event tidak ditemukan")
}
