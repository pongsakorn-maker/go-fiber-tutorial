package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	c.Send("All Books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("All Single Book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("Add a new book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Deletes a Book")
}
