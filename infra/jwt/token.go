package jwt

/*
	GenerateAccessToken 함수 / GenerateRefreshToken 함수. 만드는건 이렇게 따로 나눠놓고
	Verify는 하나로 해도 될 것 같음.


*/

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	accessExp  = jwt.NewNumericDate(time.Now().Add(time.Hour * 1))
	refreshExp = jwt.NewNumericDate(time.Now().Add(time.Hour * 120))
	secretKey  = []byte(os.Getenv("JWT_KEY")) // 이렇게 해도 되나? env 쓰지 말란 글 본적 있는데..
)

type Token struct {
	Token string
	ExpAt time.Time
}

func GenerateAccessToken(userID string, email string) (*Token, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"email":  email,
		"sub":    "ACCESS_TOKEN",
		"exp":    accessExp,
		"iat":    jwt.NewNumericDate(time.Now()),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	at := new(Token)
	at.Token = token
	at.ExpAt = accessExp.Time

	return at, nil
}

func GenerateRefreshToken(email string) (*Token, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"sub":   "REFRESH_TOKEN",
		"exp":   refreshExp,
		"iat":   jwt.NewNumericDate(time.Now()),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	rt := new(Token)
	rt.Token = token
	rt.ExpAt = refreshExp.Time

	return rt, nil
}

// access는 claims["userID"]로 내 정보 확인하거나.. (근데 헤더검증으론안쓰나)
// refresh는 로그아웃할때도, 재발급할때도 씀. 레디스 키로도 쓰이는듯(아직잘모름)
func VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Println("jwt.VerifyToken() error: %v", err)
		return nil, err
	}

	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return nil, jwt.ErrTokenExpired
	}

	return claims, nil
}
