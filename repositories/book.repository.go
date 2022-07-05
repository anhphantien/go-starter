package repositories

import (
	"go-starter/dto"
	"go-starter/entities"
	"go-starter/errors"
	"go-starter/utils"
	"net/http"

	"github.com/jinzhu/copier"
	"gorm.io/gorm/clause"
)

type BookRepository struct{}

func (repository BookRepository) FindOneByID(w http.ResponseWriter, r *http.Request, id any) (book entities.Book, err error) {
	err = CreateSqlBuilder(book).
		Joins("User").
		Where("book.id = ?", utils.ConvertToID(id)).
		Take(&book).Error
	if err != nil {
		errors.SqlError(w, r, err)
	}
	return
}

func (repository BookRepository) Create(w http.ResponseWriter, r *http.Request,
	body dto.CreateBookBody) (book entities.Book, err error) {
	copier.Copy(&book, body)
	err = CreateSqlBuilder(book).Create(&book).Error
	if err != nil {
		errors.SqlError(w, r, err)
	}
	return
}

func (repository BookRepository) Update(w http.ResponseWriter, r *http.Request, id any, body dto.UpdateBookBody) (book entities.Book, err error) {
	book, err = repository.FindOneByID(w, r, id)
	if err != nil {
		return
	}

	copier.Copy(&book, body)
	err = CreateSqlBuilder(book).
		Omit(clause.Associations). // skip auto create/update
		Updates(utils.ConvertToMap(body)).Error
	if err != nil {
		errors.SqlError(w, r, err)
	}
	return
}

func (repository BookRepository) Delete(w http.ResponseWriter, r *http.Request, id any) (err error) {
	book, err := repository.FindOneByID(w, r, id)
	if err != nil {
		return
	}

	err = CreateSqlBuilder(book).Delete(&book).Error
	if err != nil {
		errors.SqlError(w, r, err)
	}
	return
}
