package handlers

import "github.com/gofiber/fiber/v2"

type Event struct {
	Title    string
	Date     string
	Location string
	Ticket   int
}

func EventPage(c *fiber.Ctx) error {
	events := []Event{
		{
			Title:    "Konser Musik Indie",
			Date:     "20 Agustus 2025",
			Location: "Palembang",
			Ticket:   120,
		},
		{
			Title:    "Tech Conference",
			Date:     "10 September 2025",
			Location: "Jakarta",
			Ticket:   50,
		},
	}

	return c.Render("events", fiber.Map{
		"Title":  "Daftar Event",
		"Events": events,
	}, "layout")
}
