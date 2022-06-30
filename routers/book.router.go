package routers

import (
	"go-starter/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func BookRouter(r *mux.Router) {
	// r.Use(
	// 	middlewares.A,
	// 	middlewares.B,
	// )

	r.HandleFunc("/books", handlers.BookHandler{}.GetList).
		Methods(http.MethodGet)

	r.HandleFunc("/books/{id}", handlers.BookHandler{}.GetByID).
		Methods(http.MethodGet)

	r.HandleFunc("/books", handlers.BookHandler{}.Create).
		Methods(http.MethodPost)

	r.HandleFunc("/books/{id}", handlers.BookHandler{}.Update).
		Methods(http.MethodPut)
}
