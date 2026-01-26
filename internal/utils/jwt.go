package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"feeder-backend/internal/config"
)

type Claims struct {
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(email string) (string, error) {
	claims := Claims{
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := config.JWT.Secret
	return token.SignedString([]byte(secret))
}

