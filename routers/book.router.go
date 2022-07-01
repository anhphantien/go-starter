package routers

import (
	"go-starter/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

var bookHandler = handlers.BookHandler{}

func BookRouter(r *mux.Router) {
	// r.Use(
	// 	middlewares.A,
	// 	middlewares.B,
	// )

	r.HandleFunc("/books", bookHandler.GetList).
		Methods(http.MethodGet)

	r.HandleFunc("/books/{id}", bookHandler.GetByID).
		Methods(http.MethodGet)

	r.HandleFunc("/books", bookHandler.Create).
		Methods(http.MethodPost)

	r.HandleFunc("/books/{id}", bookHandler.Update).
		Methods(http.MethodPut)

	r.HandleFunc("/books/{id}", bookHandler.Delete).
		Methods(http.MethodDelete)
}
