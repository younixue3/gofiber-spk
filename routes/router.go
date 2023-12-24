package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-spk/controller"
)

func SetupRoutes(app *fiber.App) {
	//API GROUP
	api := app.Group("/api")
	api.Get("/book", controller.Books)
	api.Get("/book/:pk", controller.Book)
	api.Post("/book", controller.CreateBook)
}
