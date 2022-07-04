package middlewares

import (
	"go-starter/errors"
	"go-starter/models"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slices"
)

func RoleAuth(roles ...string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := GetCurrentUser(w, r)
			if !ok {
				return
			}
			if len(roles) > 0 && !slices.Contains(roles, user.Role) {
				errors.ForbiddenException(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
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
