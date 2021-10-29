package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pandulaDW/go-react-ambassador-app/database"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Mate 👋!")
	})

	app.Listen(":3000")
}
