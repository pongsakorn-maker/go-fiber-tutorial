package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pongsakorn-maker/go-fiber-tutorial/book"
	"github.com/pongsakorn-maker/go-fiber-tutorial/database"
)

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConnection.Close()
	setupRoutes(app)
	app.Listen(3000)
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConnection, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")
	database.DBConnection.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}
