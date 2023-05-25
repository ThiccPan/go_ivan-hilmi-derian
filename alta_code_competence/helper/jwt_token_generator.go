package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const JWT_SECRET = "secret123"

type AuthJWT interface {
	GenerateToken(email string) (string, error)
}

type authJWT struct {
	secret string
}

func NewAuthJWT() *authJWT {
	return &authJWT{
		JWT_SECRET,
	}
}

func (aj *authJWT) GenerateToken(email string) (string, error) {
	claims := &JwtCustomClaims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(JWT_SECRET))
}

type JwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}