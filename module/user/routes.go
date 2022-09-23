package user

import (
	"go-note/ent"
	"go-note/infra/middleware"
	"go-note/module/user/repository"
	"go-note/module/user/service"

	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router, db *ent.Client) {
	repo := repository.NewUserRepository(db)
	serv := service.NewUserService(repo)

	app.Get("/:username", middleware.HeaderAuthorization(), getUserHandler())
}

func getUserHandler() {

}
