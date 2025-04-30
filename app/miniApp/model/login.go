package model

import "github.com/golang-jwt/jwt/v4"

var Users = map[string]string{
	"admin": "admin",
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
}

type Claim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
