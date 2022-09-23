package auth

import (
	"testing"
)

func TestAuthRoute(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		// 첫 번째 테케
		{
			description:  "get HTTP status 200",
			route:        "/auth/register",
			expectedCode: 200,
		},
		// 두 번째 테케
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/auth/login",
			expectedCode: 404,
		},
	}

}
