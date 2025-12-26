package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"go-belajar/routes"
)

func main() {
	// template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// static file (CSS)
	app.Static("/static", "./static")

	// routes
	routes.WebRoutes(app)
	routes.ApiRoutes(app)

	// start server
	log.Fatal(app.Listen(":3000"))
}
