package routers

import (
	"go-starter/enums"
	"go-starter/handlers"
	"go-starter/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

var bookHandler = handlers.BookHandler{}

func BookRouter(r *mux.Router) {
	s := r.PathPrefix("").Subrouter()

	// s.Use(
	// 	middlewares.JwtAuth,
	// 	middlewares.RoleAuth(
	// 		enums.User.Role.ADMIN,
	// 		enums.User.Role.USER,
	// 	),
	// )

	s.HandleFunc("/books", bookHandler.GetList).
		Methods(http.MethodGet)

	s.HandleFunc("/books/{id}", bookHandler.GetByID).
		Methods(http.MethodGet)

	s.HandleFunc("/books", bookHandler.Create).
		Methods(http.MethodPost)

	s.HandleFunc("/books/{id}", bookHandler.Update).
		Methods(http.MethodPut)

	s.HandleFunc("/books/{id}",
		middlewares.NewChain(
			middlewares.JwtAuth,
			middlewares.RoleAuth(
				// enums.User.Role.ADMIN,
				enums.User.Role.USER,
			),
		).Then(
			bookHandler.Delete,
		),
	).
		Methods(http.MethodDelete)
}
