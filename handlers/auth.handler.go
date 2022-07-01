package handlers

import (
	"go-starter/dto"
	"go-starter/entities"
	"go-starter/env"
	"go-starter/errors"
	"go-starter/models"
	"go-starter/repositories"
	"go-starter/response"
	"go-starter/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct{}

// @Tags    auth
// @Summary Login
// @Param   body               body   dto.LoginBody true " "
// @Success 201                object response.Response{data=models.LoginResponse}
// @Router  /api/v1/auth/login [POST]
func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	body := dto.LoginBody{}
	if err := utils.ValidateRequestBody(w, r, &body); err != nil {
		return
	}

	user := entities.User{}
	result := repositories.
		CreateSqlBuilder(user).
		Where("username = ?", body.Username).
		Take(&user)
	if result.Error != nil {
		errors.SqlError(w, r, result.Error)
		return
	}
	if err := bcrypt.
		CompareHashAndPassword(
			[]byte(*user.HashedPassword),
			[]byte(body.Password),
		); err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			errors.BadRequestException(w, r, errors.INVALID_PASSWORD)
		default:
			errors.BadRequestException(w, r, err.Error())
		}
		return
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256,
		models.CurrentUser{
			ID:       user.ID,
			Username: *user.Username,
			Role:     *user.Role,
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().
				Add(time.Duration(env.JWT_EXPIRES_AT) * time.Second)),
		},
	).SignedString(env.JWT_SECRET)

	response.WriteJSON(w, r, response.Response{
		Data: models.LoginResponse{
			AccessToken: token,
		},
	})
}
