package repositories

import (
	"go-starter/dto"
	"go-starter/entities"
	"go-starter/errors"
	"go-starter/utils"
	"net/http"
	"sync"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookRepository struct{}

func (repository BookRepository) FindAndCount(w http.ResponseWriter, r *http.Request, q *gorm.DB) (books []entities.Book, total int64, err error) {
	ch := make(chan error, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := q.
			Session(&gorm.Session{}). // clone
			Count(&total).Error
		if err != nil {
			ch <- err
		}
	}()

	go func() {
		defer wg.Done()

		err := q.
			Session(&gorm.Session{}). // clone
			Find(&books).Error
		if err != nil {
			ch <- err
		}
	}()

	wg.Wait()
	close(ch)

	for err = range ch {
		if err != nil {
			errors.SqlError(w, r, err)
		}
	}
	return
}

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

func (repository BookRepository) Create(w http.ResponseWriter, r *http.Request, body dto.CreateBookBody) (book entities.Book, err error) {
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
		Omit(clause.Associations). // skip all associations
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
