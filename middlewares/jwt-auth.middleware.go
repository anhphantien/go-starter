package middlewares

import (
	"context"
	"go-starter/env"
	"go-starter/errors"
	"go-starter/models"
	"net/http"
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type key int

const (
	userKey key = iota
)

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	user := models.CurrentUser{}

	if r.Context().Value(userKey) == nil {
		errors.UnauthorizedException(w, r)
		return user, false
	}

	claims := r.Context().Value(userKey).(jwt.MapClaims)
	user = models.CurrentUser{
		ID:        uint64(claims["id"].(float64)),
		Username:  claims["username"].(string),
		Role:      claims["role"].(string),
		IssuedAt:  int64(claims["iat"].(float64)),
		ExpiresAt: int64(claims["exp"].(float64)),
	}
	return user, true
}
