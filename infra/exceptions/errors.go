package exceptions

import "github.com/gofiber/fiber/v2"

/*
 * global exceptions
 */

func HandleInternalServerError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"msg": "There was a problem on our side",
	})
}

func HandleInvalidInputError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"msg": err.Error(),
	})
}

func HandleUnauthorizedError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"msg": "Please login to continue",
	})
}
