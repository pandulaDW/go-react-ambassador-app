package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/pandulaDW/go-react-ambassador-app/src/database"
	"github.com/pandulaDW/go-react-ambassador-app/src/helpers"
	"github.com/pandulaDW/go-react-ambassador-app/src/middlewares"
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

// Login signs token and sends a cookie
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

// User returns the user record
func User(c *fiber.Ctx) error {
	userId, err := middlewares.GetUserId(c)
	if err != nil {
		return helpers.UnAuthorizedRequest(c, "unauthenticated")
	}

	var user models.User
	database.DB.Where("id = ?", userId).First(&user)

	return c.JSON(user)
}

// Logout logs out a user
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// UpdateInfo updates the user info
func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return helpers.BadRequest(c, "Invalid request")
	}

	id, _ := middlewares.GetUserId(c)

	user := new(models.User)
	if firstName, ok := data["first_name"]; ok {
		user.FirstName = firstName
	}
	if lastName, ok := data["last_name"]; ok {
		user.LastName = lastName
	}
	database.DB.Model(models.User{Id: id}).Updates(&user)

	var updatedUser models.User
	database.DB.Where("id = ?", id).First(&updatedUser)

	return c.JSON(updatedUser)
}

// UpdatePassword updates the user info
func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string
	fields := []string{"current_password", "new_password", "confirm_new_password"}

	if err := c.BodyParser(&data); err != nil {
		return helpers.BadRequest(c, "Invalid request")
	}

	if field, ok := helpers.RequiredFieldsIncluded(data, fields); !ok {
		return helpers.BadRequest(c, fmt.Sprintf("Missing field: '%s'", field))
	}

	id, _ := middlewares.GetUserId(c)
	var user models.User
	database.DB.Where("id = ?", id).First(&user)

	if err := user.ComparePassword([]byte(data["current_password"])); err != nil {
		return helpers.BadRequest(c, "Current password is incorrect")
	}

	if data["new_password"] != data["confirm_new_password"] {
		return helpers.BadRequest(c, "Passwords doesn't match")
	}

	user.Password = user.SetPassword([]byte(data["new_password"]))
	database.DB.Model(models.User{Id: id}).Updates(&user)

	return c.JSON(user)
}
