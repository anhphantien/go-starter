package middlewares

import (
	"context"
	"go-starter/env"
	"go-starter/errors"
	"go-starter/models"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type key int

var userKey key

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) > 6 && strings.ToUpper(tokenString[0:7]) == "BEARER " {
			tokenString = tokenString[7:]
		}

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tokenString, claims,
			func(*jwt.Token) (any, error) {
				return env.JWT_SECRET, nil
			},
		)
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

func GetCurrentUser(w http.ResponseWriter, r *http.Request) (models.CurrentUser, bool) {
	currentUser := models.CurrentUser{}

	claims, ok := r.Context().Value(userKey).(jwt.MapClaims)
	if !ok {
		errors.UnauthorizedException(w, r)
		return currentUser, false
	}

	currentUser = models.CurrentUser{
		ID:        uint64(claims["id"].(float64)),
		Username:  claims["username"].(string),
		Role:      claims["role"].(string),
		IssuedAt:  int64(claims["iat"].(float64)),
		ExpiresAt: int64(claims["exp"].(float64)),
	}
	return currentUser, true
}
