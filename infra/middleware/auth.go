package middleware

import (
	"go-note/infra/exceptions"
	"go-note/infra/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func HeaderAuthorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if len(strings.Split(authorization, " ")) < 2 {
			return exceptions.HandleUnauthorizedError(c)
		}
		accessToken := strings.Split(authorization, " ")[1]

		claims, err := jwt.VerifyToken(accessToken)
		if err != nil {
			return exceptions.HandleUnauthorizedError(c)
		}

		userID, ok := claims["userID"]
		if !ok {
			return exceptions.HandleUnauthorizedError(c)
		}

		email, ok := claims["email"]
		if !ok {
			return exceptions.HandleUnauthorizedError(c)
		}

		// 이렇게 해도 되나? 쿠키에다가..
		c.Locals("userID", userID)
		c.Locals("email", email)
		c.Locals("access_token", accessToken)

		return c.Next()
	}
}
