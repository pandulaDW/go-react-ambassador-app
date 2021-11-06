package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pandulaDW/go-react-ambassador-app/src/controllers"
)

// Setup routes of the app
func Setup(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Running smoothly!!")
	})

	admin := api.Group("/admin")
	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)
}
