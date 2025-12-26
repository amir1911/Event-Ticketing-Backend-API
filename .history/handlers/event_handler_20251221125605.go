package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Event struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Location string `json:"location"`
	Ticket   int    `json:"ticket"`
}

var events = []Event{}

func CreateEvent(c *fiber.Ctx) error {
	event := new(Event)
	if err := c.BodyParser(event); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	event.ID = len(events) + 1
	events = append(events, *event)

	return c.JSON(event)
}

func GetEvents(c *fiber.Ctx) error {
	return c.JSON(events)
}

func GetEventByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	for _, e := range events {
		if e.ID == id {
			return c.JSON(e)
		}
	}
	return c.Status(404).SendString("Event not found")
}

func UpdateEvent(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	update := new(Event)
	c.BodyParser(update)

	for i, e := range events {
		if e.ID == id {
			events[i].Title = update.Title
			events[i].Ticket = update.Ticket
			return c.JSON(events[i])
		}
	}
	return c.Status(404).SendString("Event not found")
}

func DeleteEvent(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	for i, e := range events {
		if e.ID == id {
			events = append(events[:i], events[i+1:]...)
			return c.SendString("Event deleted")
		}
	}
	return c.Status(404).SendString("Event not found")
}
