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

// SIMULASI DATABASE
var events = []Event{
	{ID: 1, Title: "Konser Musik Indie", Date: "20 Agustus 2025", Location: "Palembang", Ticket: 120},
	{ID: 2, Title: "Tech Conference", Date: "10 September 2025", Location: "Jakarta", Ticket: 50},
}

// LIST EVENT
func EventPage(c *fiber.Ctx) error {
	return c.Render("events", fiber.Map{
		"Title":  "Daftar Event",
		"Events": events,
	})
}

// DETAIL EVENT
func EventDetail(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendString("ID tidak valid")
	}

	for _, e := range events {
		if e.ID == id {
			return c.Render("detail", fiber.Map{
				"Event": e,
			})
		}
	}

	return c.SendString("Event tidak ditemukan")
}

// ✅ BOOKING + GATEKEEPER LOGIC (INTI)
func BookTicket(c *fiber.Ctx) error {
	eventID, _ := strconv.Atoi(c.FormValue("event_id"))
	qty, _ := strconv.Atoi(c.FormValue("qty"))

	// validasi dasar
	if qty <= 0 {
		return c.Render("result", fiber.Map{
			"Message": "❌ Jumlah tiket tidak valid",
		})
	}

	for i, e := range events {
		if e.ID == eventID {

			// 🔒 GATEKEEPER LOGIC
			if qty > e.Ticket {
				return c.Render("result", fiber.Map{
					"Message": "❌ Tiket tidak cukup (Sold Out)",
				})
			}

			// kurangi stok (TRANSAKSI BERHASIL)
			events[i].Ticket -= qty

			return c.Render("result", fiber.Map{
				"Message": "✅ Booking berhasil! Sisa tiket: " + strconv.Itoa(events[i].Ticket),
			})
		}
	}

	return c.Render("result", fiber.Map{
		"Message": "Event tidak ditemukan",
	})
}
