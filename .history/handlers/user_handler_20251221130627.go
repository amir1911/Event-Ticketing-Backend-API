package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-belajar/handlers"
)

func ApiRoutes(app *fiber.App) {
	api := app.Group("/api")

	// ======================
	// EVENT CRUD
	// ======================
	api.Post("/events", handlers.CreateEvent)
	api.Get("/events", handlers.GetEvents)
	api.Get("/events/:id", handlers.GetEventByID)
	api.Put("/events/:id", handlers.UpdateEvent)
	api.Delete("/events/:id", handlers.DeleteEvent)

	// ======================
	// USER CRUD
	// ======================
	api.Post("/users", handlers.CreateUser)
	api.Get("/users", handlers.GetUsers)        // opsional tapi bagus
	api.Get("/users/:id", handlers.GetUserByID)
	api.Put("/users/:id", handlers.UpdateUser)
	api.Delete("/users/:id", handlers.DeleteUser)

	// ======================
	// BOOKING / TRANSACTION
	// ======================
	api.Post("/bookings", handlers.CreateBooking)
	api.Get("/bookings", handlers.GetBookings)
}
