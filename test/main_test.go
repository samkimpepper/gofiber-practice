package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		// 첫 번째 테케
		{
			description:  "get HTTP status 200",
			route:        "/test",
			expectedCode: 200,
		},
		// 두 번째 테케
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/not-found",
			expectedCode: 404,
		},
	}

	app := fiber.New()

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)

		resp, _ := app.Test(req, 1)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}

}
