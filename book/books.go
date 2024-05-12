package book

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json: "id"`
	Title  string `json: "title"`
	Author string `json: "author"`
}

var books []Book

func StartBooks() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "Title1", Author: "Nun"})
	books = append(books, Book{ID: 2, Title: "Title2", Author: "Nun2"})
	books = append(books, Book{ID: 3, Title: "Title3", Author: "Nun3"})

	app.Get("/books", getBooks)
	app.Get("/book/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("books/:id", deleteBook)

	app.Listen(":8080")
}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	bookId, error := strconv.Atoi(c.Params("id"))

	if error != nil {
		c.Status(fiber.StatusBadRequest).SendString(error.Error())
	}

	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)
	if error := c.BodyParser(book); error != nil {
		return c.Status(fiber.StatusBadRequest).SendString(error.Error())
	}

	for _, tempBook := range books {
		if tempBook.ID == book.ID {
			return c.Status(fiber.StatusFound).SendString("User Id Duplicate")
		}
	}

	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Parser Book update
	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for index, book := range books {
		if book.ID == bookId {
			books[index].Title = bookUpdate.Title
			books[index].Author = bookUpdate.Author
			return c.JSON(books[index])
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func deleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
