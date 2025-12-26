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

var events = []Event{
	{ID: 1, Title: "Konser Musik Indie", Date: "20 Agustus 2025", Location: "Palembang", Ticket: 120},
	{ID: 2, Title: "Tech Conference", Date: "10 September 2025", Location: "Jakarta", Ticket: 50},
}

func EventPage(c *fiber.Ctx) error {
	return c.Render("events", fiber.Map{
		"Title":  "Daftar Event",
		"Events": events,
	})
}

func EventDetail(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendString("ID tidak valid")
	}

	for _, e := range events {
		if e.ID == id {
			return c.Render("detail", fiber.Map{
				"Title": "Detail Event",
				"Event": e,
			})
		}
	}

	return c.SendString("Event tidak ditemukan")
}
