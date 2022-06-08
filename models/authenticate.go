package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Authenticate struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type IDTokenClaims struct {
	jwt.StandardClaims
	UserID string `json:"userID"`
}
