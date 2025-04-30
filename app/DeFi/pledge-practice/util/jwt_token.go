package util

import (
	"github.com/dgrijalva/jwt-go"
	"p-cy/config"
	"p-cy/log"
	"time"
)

func CreateToken(username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	signedString, err := claims.SignedString([]byte(config.GlobalConfig.Jwt.Secret))
	if err != nil {
		log.Logger.Error("create jwt token fail")
	}
	return signedString, nil
}

func ParseToken(token string) (string, error) {
	parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GlobalConfig.Jwt.Secret), nil
	})
	if err != nil || !parse.Valid {
		log.Logger.Error("parse jwt token fail")
	}
	username := parse.Claims.(jwt.MapClaims)["username"].(string)
	return username, nil

}
