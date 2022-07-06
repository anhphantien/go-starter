package repositories

import (
	"go-starter/entities"
	"go-starter/errors"
	"go-starter/utils"
	"net/http"
)

type UserRepository struct{}

func (repository UserRepository) FindByID(w http.ResponseWriter, r *http.Request, id any) (user entities.User, err error) {
	err = CreateSqlBuilder(user).
		Where("id = ?", utils.ConvertToID(id)).
		Take(&user).Error
	if err != nil {
		errors.SqlError(w, r, err)
	}
	return
}
