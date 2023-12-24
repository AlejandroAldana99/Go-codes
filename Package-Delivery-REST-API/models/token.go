package models

import "github.com/golang-jwt/jwt/v4"

type TokenData struct {
	UserID         string
	Name           string
	Role           string
	StandardClaims *jwt.Claims
}
