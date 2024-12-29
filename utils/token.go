package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateToken(id uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
