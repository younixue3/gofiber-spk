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
	if err := ctx.BodyParser(&book); err != nil {
		return err
	}
	database.Db.Create(book)
	return ctx.JSON(book)
}
