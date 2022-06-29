package routers

import (
	"go-starter/handlers"
	"go-starter/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router) {
	// r.Use(
	// 	middlewares.A,
	// 	middlewares.B,
	// )

	r.HandleFunc(
		"/auth/login",
		middlewares.NewChain(
			middlewares.A,
			middlewares.B,
		).Then(
			handlers.AuthHandler{}.Login,
		),
	).Methods(http.MethodPost)

	r.HandleFunc(
		"/auth/login",
		handlers.AuthHandler{}.Login,
	).Methods(http.MethodPost)
}
