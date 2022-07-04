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
	r.Use(
		middlewares.JwtAuth,
		middlewares.RoleAuth(
			enums.User.Role.ADMIN,
			enums.User.Role.USER,
		),
	)

	r.HandleFunc("/books", bookHandler.GetList).
		Methods(http.MethodGet)

	r.HandleFunc("/books/{id}", bookHandler.GetByID).
		Methods(http.MethodGet)

	r.HandleFunc("/books", bookHandler.Create).
		Methods(http.MethodPost)

	r.HandleFunc("/books/{id}", bookHandler.Update).
		Methods(http.MethodPut)

	r.HandleFunc("/books/{id}",
		// middlewares.NewChain(
		// 	middlewares.JwtAuth,
		// 	middlewares.UserRoles(
		// 		enums.User.Role.ADMIN,
		// 		enums.User.Role.USER,
		// 	),
		// ).Then(
		// 	bookHandler.Delete,
		// ),
		bookHandler.Delete,
	).
		Methods(http.MethodDelete)
}
