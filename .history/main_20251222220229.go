package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go-belajar/config"
	"go-belajar/database"
	"go-belajar/routes"
)

func main() {
	config.ENVLoad()
	app := fiber.New()


	database.DBMigrate()

	routes.ApiRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
