package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-belajar/handlers"
)

func WebRoutes(app *fiber.App) {
	// LIST EVENT
	app.Get("/events", handlers.EventPage)

	// DETAIL EVENT (INI PENTING)
	app.Get("/events/:id", handlers.EventDetail)

	// BOOKING
	app.Post("/book", handlers.BookTicket)
}
