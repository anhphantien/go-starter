package utils

import (
	"go-starter/env"
	"go-starter/errors"
	"go-starter/models"
	"net/http"
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func GetCurrentUser(w http.ResponseWriter, r *http.Request) (models.CurrentUser, error) {
	user := models.CurrentUser{}

	claims, err := verifyToken(r)
	if err != nil {
		switch strings.ToLower(err.Error()) {
		case jwt.ErrTokenExpired.Error():
			errors.UnauthorizedException(w, r, jwt.ErrTokenExpired.Error())
		default:
			errors.UnauthorizedException(w, r, jwt.ErrTokenMalformed.Error())
		}
		return user, err
	}

	user = models.CurrentUser{
		ID:        uint64(claims["id"].(float64)),
		Username:  claims["username"].(string),
		Role:      claims["role"].(string),
		IssuedAt:  claims["iat"].(*jwt.NumericDate),
		ExpiresAt: claims["exp"].(*jwt.NumericDate),
	}
	// if ok := validateUserRole(c, user); !ok {
	// 	return user, errors.ForbiddenException(c), false
	// }
	return user, err
}

// func validateUserRole(c *fiber.Ctx, user models.CurrentUser) bool {
// 	roles := []string{}

// 	if ADMIN, ok := c.Locals(middlewares.ADMIN_ROLE).(string); ok {
// 		roles = append(roles, ADMIN)
// 	}
// 	if USER, ok := c.Locals(middlewares.USER_ROLE).(string); ok {
// 		roles = append(roles, USER)
// 	}

// 	if len(roles) > 0 && !slices.Contains(roles, user.Role) {
// 		return false
// 	}
// 	return true
// }

func verifyToken(r *http.Request) (jwt.MapClaims, error) {
	tokenString := r.Header.Get("Authorization")
	if strings.HasPrefix(tokenString, "Bearer ") || strings.HasPrefix(tokenString, "bearer ") {
		tokenString = regexp.MustCompile(`[B|b]earer\s+`).ReplaceAllString(tokenString, "")
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(env.JWT_SECRET), nil
		},
	)

	return claims, err
}
