package models

import "github.com/golang-jwt/jwt/v4"

type CurrentUser struct {
	ID        uint64           `json:"id"`
	Username  string           `json:"username"`
	Role      string           `json:"role"`
	IssuedAt  *jwt.NumericDate `json:"iat"`
	ExpiresAt *jwt.NumericDate `json:"exp"`
	jwt.RegisteredClaims
}
