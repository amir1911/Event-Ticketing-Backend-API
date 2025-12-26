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

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	update := new(User)
	c.BodyParser(update)

	for i, u := range users {
		if u.ID == id {
			users[i].Name = update.Name
			users[i].Email = update.Email
			users[i].Role = update.Role
			return c.JSON(users[i])
		}
	}

	return c.Status(404).SendString("User not found")
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return c.SendString("User deleted")
		}
	}

	return c.Status(404).SendString("User not found")
}
