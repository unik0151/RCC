package service

import (
	"github.com/golang-jwt/jwt/v4"
	"miniapp/model"
	"time"
)

var (
	expireTime = time.Hour * 24
	jwtSecret  = []byte("2b129834fe16c944133")
)

func GenerateToken(username string) (string, error) {
	claims := &model.Claim{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "miniapp",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GetJwtKey() []byte {
	return jwtSecret
}
