package handlers

import (
	"time" // 1. Import package time

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your-secret-key")

func CreateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["iat"] = time.Now().Unix()                      // Waktu token dibuat (detik saat ini)
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()  // Token akan kadaluarsa dalam 24 jam

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
