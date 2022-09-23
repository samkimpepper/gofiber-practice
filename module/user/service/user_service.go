package service

import (
	"go-note/module/user/repository"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

type UserService interface {
	GetUser(c *fiber.Ctx) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// ========================================================

func (s userService) GetUser(c *fiber.Ctx) error {
	username := c.Params("username")
	userID := c.Locals("userID")

	uuid := uuid.Parse(userID)

}
