package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-belajar/controllers"
)

func ApiRoutes(app *fiber.App) {
	api := app.Group("/api")

	// USER
	api.Post("/users", handlers.CreateUser)
	api.Get("/users", handlers.GetUsers)
	api.Get("/users/:id", handlers.GetUserByID)
	api.Put("/users/:id", handlers.UpdateUser)
	api.Delete("/users/:id", handlers.DeleteUser)

	// EVENT
	api.Post("/events", handlers.CreateEvent)
	api.Get("/events", handlers.GetEvents)
	api.Get("/events/:id", handlers.GetEventByID)
	api.Put("/events/:id", handlers.UpdateEvent)
	api.Delete("/events/:id", handlers.DeleteEvent)

	// BOOKING
	api.Post("/bookings", handlers.CreateBooking)
	api.Get("/bookings", handlers.GetBookings)
}
