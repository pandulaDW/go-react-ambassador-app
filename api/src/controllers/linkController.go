package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pandulaDW/go-react-ambassador-app/src/database"
	"github.com/pandulaDW/go-react-ambassador-app/src/helpers"
	"github.com/pandulaDW/go-react-ambassador-app/src/models"
	"strconv"
)

func Link(c *fiber.Ctx) error {
	id := c.Params("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return helpers.BadRequest(c, "Invalid user id")
	}

	user := new(models.User)
	database.DB.Where("id = ?", userId).First(user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{"message": "No user found with the given Id"})
	}

	var links []models.Link
	database.DB.Where("user_id = ?", userId).Find(&links)

	return c.JSON(links)
}
