package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-spk/models"
	"go-fiber-spk/settings/database"
	"log"
)

var book models.Book

func Books(ctx *fiber.Ctx) error {
	results := database.Db.Take(&book)
	if results.Error != nil {
		log.Println(results.Error)
	}
	return ctx.JSON(book)
}

func Book(ctx *fiber.Ctx) error {
	pk := ctx.Params("pk")
	results := database.Db.First(&book, pk)
	if results.Error != nil {
		log.Println(results.Error)
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    book,
	})
}

func CreateBook(ctx *fiber.Ctx) error {
	err := ctx.BodyParser(&book)
	if err != nil {
		return err
	}
	recordBook := models.Book{
		Title: book.Title,
	}

	createBook := database.Db.Create(&book)
	if createBook.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "Failed to store data"})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    recordBook,
	})
}

func UpdateBook(ctx *fiber.Ctx) error {
	pk := ctx.Params("pk")

	recordBook := models.Book{
		Title: book.Title,
	}

	err := ctx.BodyParser(&book)
	if err != nil {
		return err
	}

	updateBook := database.Db.Model(&book).Where("id = ?", pk).Updates(&book)
	if updateBook.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "Failed to store data"})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    recordBook,
	})
}

func DeleteBook(ctx *fiber.Ctx) error {
	pk := ctx.Params("pk")

	recordBook := database.Db.First(&book, pk)
	if recordBook.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "Data not found"})
	}

	deleteBook := database.Db.Delete(&book)
	if deleteBook.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "Failed to delete data"})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}
