package models

import "github.com/golang-jwt/jwt/v4"

type CurrentUser struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
	jwt.RegisteredClaims
}
