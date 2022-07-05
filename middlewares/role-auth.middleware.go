package middlewares

import (
	"go-starter/errors"
	"net/http"

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
