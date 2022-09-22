package service

import (
	"errors"
	"go-note/infra/exceptions"
	"go-note/module/auth/entities"
	"go-note/module/auth/repository"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Register(c *fiber.Ctx) error
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

// ========================================================

// register
func (s authService) Register(c *fiber.Ctx) error {
	var dto entities.RegisterRequest
	if err := c.BodyParser(&dto); err != nil {
		return exceptions.HandleInvalidInputError(c, err)
	}

	if err := hashPassword(&dto.Password); err != nil {
		log.Println("authService.hashPassword error: %v", err)
		return exceptions.HandleInternalServerError(c)
	}

	_, err := s.repo.Save(dto)
	if err != nil {
		if errors.Is(err, entities.ErrUserAlreadyExists) {
			return entities.HandleUserAlreadyExists(c)
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Successfully registered a new account",
	})
}

// bcrypt
func hashPassword(pw *string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*pw), 8)
	*pw = string(bytes)
	return err
}
func verifyPassword(pw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	return err == nil
}
