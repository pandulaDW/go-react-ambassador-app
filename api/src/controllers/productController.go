package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pandulaDW/go-react-ambassador-app/src/database"
	"github.com/pandulaDW/go-react-ambassador-app/src/helpers"
	"github.com/pandulaDW/go-react-ambassador-app/src/models"
	"net/http"
	"strconv"
)

// Products returns all the products
func Products(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)

	return c.JSON(products)
}

// CreateProduct creates a new product
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return helpers.BadRequest(c, "Invalid request")
	}

	database.DB.Create(&product)

	return c.JSON(product)
}

// GetProduct returns a particular product
func GetProduct(c *fiber.Ctx) error {
	var product models.Product
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helpers.BadRequest(c, "Invalid value for id")
	}

	database.DB.Where("id = ?", productId).Find(&product)
	return c.JSON(product)
}

// UpdateProduct returns a particular product
func UpdateProduct(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helpers.BadRequest(c, "Invalid value for id")
	}

	product := models.Product{Id: uint(productId)}
	if err = c.BodyParser(&product); err != nil {
		return helpers.BadRequest(c, "Invalid request")
	}

	database.DB.Model(&product).Updates(&product)
	return c.JSON(product)
}

// DeleteProduct returns a particular product
func DeleteProduct(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helpers.BadRequest(c, "Invalid value for id")
	}

	product := models.Product{Id: uint(productId)}

	database.DB.Delete(&product)
	return c.Status(http.StatusNoContent).JSON("")
}
