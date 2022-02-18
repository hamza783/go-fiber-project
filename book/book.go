package book

import (
	"github.com/gofiber/fiber"
	db "github.com/hamza/go-fiber-tutorial/database"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	var books []Book
	db.DB.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	var book Book
	db.DB.Find(&book, id)
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.DB.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	var book Book
	db.DB.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No book found with id")
		return
	}
	db.DB.Delete(&book)
	c.Status(200).Send("Book successfully deleted")
}
