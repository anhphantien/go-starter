package repositories

import (
	"go-starter/entities"
	"go-starter/utils"
)

type UserRepository struct{}

func (r UserRepository) FindOneByID(id any) (user entities.User, err error) {
	err = CreateSqlBuilder(user).
		Where("id = ?", utils.ConvertToID(id)).
		Take(&user).Error
	return user, err
}
