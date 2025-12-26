package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-belajar/controllers"
)

func ApiRoutes(app *fiber.App) {
	api := app.Group("/api")

	// USER
	api.Post("/users", controllers.CreateUser)
	api.Get("/users", controllers.GetUsers)
	api.Get("/users/:id", controllers.GetUserByID)
	api.Put("/users/:id", controllers.UpdateUser)
	api.Delete("/users/:id", controllers.DeleteUser)

	// EVENT
	api.Post("/events", controllers.CreateEvent)
	api.Get("/events", controllers.GetEvents)
	api.Get("/events/:id", controllers.GetEventByID)
	api.Put("/events/:id", controllers.UpdateEvent)
	api.Delete("/events/:id", controllers.DeleteEvent)

	// BOOKING
	api.Post("/bookings", controllers.CreateBooking)
	api.Get("/bookings", controllers.GetBookings)
}
