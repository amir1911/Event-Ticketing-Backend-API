package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-belajar/handlers"
)

func WebRoutes(app *fiber.App) {
	app.Get("/events", handlers.EventPage)
}
