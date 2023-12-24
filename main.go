package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-spk/routes"
	"go-fiber-spk/settings/database"
)

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
	database.ConnectDatabase()
	database.RunMigration()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
