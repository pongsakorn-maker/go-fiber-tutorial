package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/pongsakorn-maker/go-fiber-tutorial/database"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConnection
	var books []Book
	db.Find(&books)
	c.JSON(books)
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
