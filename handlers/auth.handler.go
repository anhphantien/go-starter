package handlers

import (
	"fmt"
	"go-starter/common"
	"go-starter/dto"
	"go-starter/entities"
	"go-starter/repositories"
	"go-starter/utils"
	"net/http"
)

type AuthHandler struct{}

// @Tags    auth
// @Summary Login
// @Param   body               body     dto.LoginBody true " "
// @Success 201                {object} string
// @Router  /api/v1/auth/login [POST]
func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	body := dto.LoginBody{}
	if err := utils.ValidateRequestBody(r, &body); err != nil {
		http.Error(w, "failed to validate struct", 400)
		return
	}

	fmt.Println(2222222222222222, body.Password, body.Username)

	user := entities.User{}
	res := repositories.
		CreateSqlBuilder(user).
		Where("username = ?", body.Username).
		Take(&user)
	fmt.Println(3333333333333333, res)

	common.WriteJSON(w, http.StatusOK, "ối dồi ôi")
}
