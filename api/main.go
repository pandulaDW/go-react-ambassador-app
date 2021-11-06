package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pandulaDW/go-react-ambassador-app/src/database"
	"github.com/pandulaDW/go-react-ambassador-app/src/routes"
	"log"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowCredentials: true}))

	routes.Setup(app)

	log.Fatal(app.Listen(":8000"))
}
