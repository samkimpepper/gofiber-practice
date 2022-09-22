package entities

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrExpiredToken      = errors.New("token expired")
)

func HandleUserAlreadyExists(c *fiber.Ctx) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"msg": "User already exists",
	})
}
