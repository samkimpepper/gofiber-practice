package auth

import (
	"go-note/ent"
	"go-note/module/auth/repository"
	"go-note/module/auth/service"

	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router, db *ent.Client) {
	repo := repository.NewAuthRepository(db)
	serv := service.NewAuthService(repo)

	app.Post("/register", registerHandler(serv))
}

func registerHandler(serv service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return serv.Register(c)
	}
}
