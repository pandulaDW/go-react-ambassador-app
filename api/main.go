package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pandulaDW/go-react-ambassador-app/src/database"
	"github.com/pandulaDW/go-react-ambassador-app/src/routes"
	"log"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()
	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
