package middlewares

import (
	"context"
	"go-starter/env"
	"go-starter/errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type key int

var userKey key

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := verifyToken(r)
		if err != nil {
			switch strings.ToLower(err.Error()) {
			case jwt.ErrTokenExpired.Error():
				errors.UnauthorizedException(w, r, jwt.ErrTokenExpired.Error())
			default:
				errors.UnauthorizedException(w, r, jwt.ErrTokenMalformed.Error())
			}
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(),
			userKey, claims,
		)))
	})
}

func verifyToken(r *http.Request) (jwt.MapClaims, error) {
	tokenString := r.Header.Get("Authorization")
	if strings.HasPrefix(tokenString, "Bearer ") || strings.HasPrefix(tokenString, "bearer ") {
		tokenString = regexp.MustCompile(`[B|b]earer\s+`).ReplaceAllString(tokenString, "")
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (any, error) {
			return []byte(env.JWT_SECRET), nil
		},
	)

	return claims, err
}
