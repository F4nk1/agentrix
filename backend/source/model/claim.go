package model

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserId	int 	`json:"user_id"`
	RoleId	string	`json:"role_id"`
	jwt.RegisteredClaims
}