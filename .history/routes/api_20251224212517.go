package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-belajar/controllers"
	"go-belajar/middleware"
)

func ApiRoutes(app *fiber.App) {
	api := app.Group("/api")

	// =====================
	// AUTH
	// =====================
	api.Post("/login", controllers.Login)

	// =====================
	// USER
	// =====================
	api.Post("/users", controllers.CreateUser)
	api.Get("/users", controllers.GetUsers)
	api.Get("/users/:id", controllers.GetUserByID)
	api.Put("/users/:id", controllers.UpdateUser)
	api.Delete("/users/:id", controllers.DeleteUser)

	// =====================
	// EVENT (ADMIN ONLY)
	// =====================
	api.Post(
		"/events",
		middleware.AuthRequired,
		middleware.AdminOnly,
		controllers.CreateEvent,
	)
	api.Get("/events", controllers.GetEvents)
	api.Get("/events/:id", controllers.GetEventByID)
	api.Put(
		"/events/:id",
		middleware.AuthRequired,
		middleware.AdminOnly,
		controllers.UpdateEvent,
	)
	api.Delete(
		"/events/:id",
		middleware.AuthRequired,
		middleware.AdminOnly,
		controllers.DeleteEvent,
	)

	// =====================
	// BOOKING (LOGIN ONLY)
	// =====================
	api.Post(
		"/bookings",
		middleware.AuthRequired,
		controllers.CreateBooking,
	)
	api.Get(
		"/bookings",
		middleware.AuthRequired,
		controllers.GetBookings,
	)
}
