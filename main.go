package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/hamza/go-fiber-tutorial/book"
	db "github.com/hamza/go-fiber-tutorial/database"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello Word")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	if err := db.Open(); err != nil {
		// handle error
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	db.DB.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer db.Close()

	setupRoutes(app)

	app.Listen(3000)
}
