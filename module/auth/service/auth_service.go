package service

import (
	"errors"
	"go-note/infra/exceptions"
	"go-note/infra/jwt"
	"go-note/module/auth/entities"
	"go-note/module/auth/repository"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
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

// login
func (s authService) Login(c *fiber.Ctx) error {
	var dto entities.LoginRequest
	if err := c.BodyParser(&dto); err != nil {
		return exceptions.HandleInvalidInputError(c, err)
	}

	user, err := s.repo.FindByEmail(dto.Email)
	if err != nil {
		return entities.HandleUserDoesNotExist(c)
	}

	if !verifyPassword(dto.Password, user.Password) {
		return entities.HandleInvalidPassword(c) // 이걸 따로 구분해줘도 되나
	}

	// 토큰
	at, err := jwt.GenerateAccessToken(user.ID.String(), user.Email)
	if err != nil {
		log.Println("service.authService.Login(): %v", err)
		return exceptions.HandleInternalServerError(c)
	}

	rt, err := jwt.GenerateRefreshToken(user.Email)
	if err != nil {
		log.Println("service.authService.Login(): %v", err)
		return exceptions.HandleInternalServerError(c)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    rt.Token,
		Expires:  rt.ExpAt,
		HTTPOnly: true,
		Path:     "/",
		SameSite: "None",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": at.Token,
	})
}

// func (s authService) Logout(c *fiber.Ctx) error {
// 	rt := c.Cookies("refresh_token")
// 	if rt == "" {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"msg": "Please login to continue",
// 		})
// 	}

// }
