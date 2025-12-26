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
	database.DBLoad()    // ⬅️ HARUS sukses
	database.DBMigrate() // ⬅️ BARU migrate

	app := fiber.New()
	routes.ApiRoutes(app)

	app.Listen(":3000")
}
	
