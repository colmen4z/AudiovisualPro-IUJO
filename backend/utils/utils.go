package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"audiovisualpro/backend/models"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(gestorID uint, usuario string) (string, error) {
	claims := &models.GestorClaims{
		ID: gestorID,
		Usuario: usuario,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour*24)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "audiovisualpro-backend",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}