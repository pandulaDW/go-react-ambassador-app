package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/pandulaDW/go-react-ambassador-app/src/database"
	"github.com/pandulaDW/go-react-ambassador-app/src/helpers"
	"github.com/pandulaDW/go-react-ambassador-app/src/models"
	"strconv"
	"time"
)

// Register handler registers a new user
func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return helpers.BadRequest(c, "Invalid request")
	}

	if data["password"] != data["password_confirm"] {
		return helpers.BadRequest(c, "passwords do not match")
	}

	user := models.User{
		FirstName:    data["first_name"],
		LastName:     data["last_name"],
		Email:        data["email"],
		IsAmbassador: false,
	}
	user.Password = user.SetPassword([]byte(data["password"]))

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return helpers.BadRequest(c, "Invalid request")
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		return helpers.BadRequest(c, "User not found")
	}

	if err := user.ComparePassword([]byte(data["password"])); err != nil {
		return helpers.BadRequest(c, "Invalid credentials")
	}

	payload := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Subject:   strconv.Itoa(int(user.Id)),
	}

	token, tokenErr := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))
	if tokenErr != nil {
		return helpers.BadRequest(c, "Invalid credentials")
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
