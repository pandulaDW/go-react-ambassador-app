package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pandulaDW/go-react-ambassador-app/src/controllers"
	"github.com/pandulaDW/go-react-ambassador-app/src/middlewares"
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

	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Post("/logout", controllers.Logout)
	adminAuthenticated.Get("/user", controllers.User)
	adminAuthenticated.Put("/user/info", controllers.UpdateInfo)
	adminAuthenticated.Put("/user/password", controllers.UpdatePassword)
	adminAuthenticated.Get("/ambassadors", controllers.Ambassador)
}
