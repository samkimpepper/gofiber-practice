package handler

import (
	"go-note/module/auth/service"

	"github.com/gofiber/fiber"
)

// register, login, reissue, logout

// register
func Register(service service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {

	}
}
