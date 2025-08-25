package models

import "github.com/golang-jwt/jwt/v5"

type JwtToken struct {
	UserID   int `json:"id"`
	UserRole int `json:"role"`
	jwt.RegisteredClaims
}

type Token struct {
	Token string `json:"token"`
}
