package entities

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserDoesNotExist  = errors.New("user does not exist")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrExpiredToken      = errors.New("token expired")
)

func HandleUserAlreadyExists(c *fiber.Ctx) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"msg": "User already exists",
	})
}

func HandleUserDoesNotExist(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"msg": "User does not exist", // 이거 그냥 user/pw 틀렸다고만 알려줘야되나?
	})
}

func HandleInvalidPassword(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"msg": "User does not exist", // 이거 그냥 user/pw 틀렸다고만 알려줘야되나?
	})
}
