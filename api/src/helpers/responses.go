package helpers

import "github.com/gofiber/fiber/v2"

// BadRequest sends a http bad request
func BadRequest(c *fiber.Ctx, message string) error {
	c.Status(fiber.StatusBadRequest)
	return c.JSON(fiber.Map{"message": message})
}

// UnAuthorizedRequest sends a http bad request
func UnAuthorizedRequest(c *fiber.Ctx, message string) error {
	c.Status(fiber.StatusUnauthorized)
	return c.JSON(fiber.Map{"message": message})
}
