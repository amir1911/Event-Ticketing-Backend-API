package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go-belajar/config"
	"go-belajar/database"
	"go-belajar/routes"
)

func main() {
	app := fiber.New()

	config.ConnectDB()
	database.Migrate()

	routes.ApiRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
