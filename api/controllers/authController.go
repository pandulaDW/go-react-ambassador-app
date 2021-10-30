package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pandulaDW/go-react-ambassador-app/database"
	"github.com/pandulaDW/go-react-ambassador-app/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Register handler registers a new user
func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

	user := models.User{
		FirstName:    data["first_name"],
		LastName:     data["last_name"],
		Email:        data["email"],
		Password:     encryptedPassword,
		IsAmbassador: false,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}
