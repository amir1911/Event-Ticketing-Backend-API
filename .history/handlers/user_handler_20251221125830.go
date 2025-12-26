package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"` // organizer / customer
}

var users = []User{}

func CreateUser(c *fiber.Ctx) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	user.ID = len(users) + 1
	users = append(users, *user)

	return c.JSON(user)
}

func GetUserByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	for _, u := range users {
		if u.ID == id {
			return c.JSON(u)
		}
	}

	return c.Status(404).SendString("User not found")
}
