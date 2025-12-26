package main

import (
	"github.com/gofiber/fiber/v2"

	"go-belajar/config"
	"go-belajar/database"
	"go-belajar/routes"
)

func main() {
	// 1️⃣ Load ENV
	config.ENVLoad()

	// 2️⃣ Connect Database
	database.ConnectDB()

	// 3️⃣ Auto Migrate
	database.DBMigrate()

	// 4️⃣ Start Fiber
	app := fiber.New()
	routes.ApiRoutes(app)

	app.Listen(":8080")
}
