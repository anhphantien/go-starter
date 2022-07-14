package handlers

import (
	"database/sql"
	"go-starter/dto"
	"go-starter/entities"
	"go-starter/models"
	"go-starter/repositories"
	"go-starter/response"
	"go-starter/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type BookHandler struct{}

// @Tags    books
// @Summary Get a list of books
// @Param   limit         query  int    false " "
// @Param   page          query  int    false " "
// @Param   keyword       query  string false " "
// @Param   filter        query  object false " "
// @Param   sort          query  object false " "
// @Success 200           object response.Response{data=[]entities.Book}
// @Router  /api/v1/books [GET]
func (h BookHandler) GetList(w http.ResponseWriter, r *http.Request) {
	pagination := utils.Pagination(r)

	books := []entities.Book{}
	q := repositories.CreateSqlBuilder(books).
		Preload("User")
	if pagination.Filter["id"] != nil {
		q.Where("book.id = ?", utils.ConvertToID(pagination.Filter["id"]))
	}
	if len(pagination.Keyword) > 0 {
		q.Where(
			"book.title LIKE @keyword OR book.description LIKE @keyword",
			sql.Named("keyword", "%"+pagination.Keyword+"%"),
		)
	}
	q.Limit(pagination.Limit).
		Offset(pagination.Offset).
		Order(pagination.Order)
	books, total, err := bookRepository.FindAndCount(w, r, q)
	if err != nil {
		return
	}

	// var err error

	// var total int64
	// err = q.Count(&total).Error
	// if err != nil {
	// 	errors.SqlError(w, r, err)
	// 	return
	// }

	// err = q.
	// 	Limit(pagination.Limit).
	// 	Offset(pagination.Offset).
	// 	Order(pagination.Order).
	// 	Find(&books).Error
	// if err != nil {
	// 	errors.SqlError(w, r, err)
	// 	return
	// }

	response.WriteJSON(w, r, response.Response{
		Data: models.PaginationResponse{
			Items: books,
			Total: total,
		},
	})
}

// @Tags    books
// @Summary Get a book by ID
// @Param   id                 path   int true " "
// @Success 200                object response.Response{data=entities.Book}
// @Router  /api/v1/books/{id} [GET]
func (h BookHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	book, err := bookRepository.FindOneByID(w, r, id)
	if err != nil {
		return
	}

	response.WriteJSON(w, r, response.Response{
		Data: book,
	})
}

// @Tags    books
// @Summary Create a new book
// @Param   body          body   dto.CreateBookBody true " "
// @Success 201           object response.Response{data=entities.Book}
// @Router  /api/v1/books [POST]
func (h BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	body := dto.CreateBookBody{}
	if err := utils.ValidateRequestBody(w, r, &body); err != nil {
		return
	}

	if body.UserID != nil {
		_, err := userRepository.FindOneByID(w, r, body.UserID)
		if err != nil {
			return
		}
	}

	book, err := bookRepository.Create(w, r, body)
	if err != nil {
		return
	}

	response.WriteJSON(w, r, response.Response{
		Data: book,
	})
}

// @Tags    books
// @Summary Update a book
// @Param   id                 path   int true " "
// @Param   body               body   dto.UpdateBookBody true " "
// @Success 200                object response.Response{data=entities.Book}
// @Router  /api/v1/books/{id} [PUT]
func (h BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	body := dto.UpdateBookBody{}
	if err := utils.ValidateRequestBody(w, r, &body); err != nil {
		return
	}

	id := mux.Vars(r)["id"]

	book, err := bookRepository.Update(w, r, id, body)
	if err != nil {
		return
	}

	response.WriteJSON(w, r, response.Response{
		Data: book,
	})
}

// @Security Bearer
// @Summary  Delete a book
// @Tags     books
// @Param    id                 path     int true " "
// @Success  200                object   response.Response{data=boolean}
// @Router   /api/v1/books/{id} [DELETE]
func (h BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// currentUser, ok := middlewares.GetCurrentUser(w, r)
	// if !ok {
	// 	return
	// }

	id := mux.Vars(r)["id"]

	err := bookRepository.Delete(w, r, id)
	if err != nil {
		return
	}

	response.WriteJSON(w, r, response.Response{
		Data: true,
	})
}
