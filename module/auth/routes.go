package auth

import (
	"go-note/ent"
	"go-note/module/auth/repository"
	"go-note/module/auth/service"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router, db *ent.Client, rdb *redis.Client) {
	repo := repository.NewAuthRepository(db, rdb)
	serv := service.NewAuthService(repo)

	app.Post("/register", registerHandler(serv))
	app.Post("/login", loginHandler(serv))
}

func registerHandler(serv service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return serv.Register(c)
	}
}

func loginHandler(serv service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return serv.Login(c)
	}
}
