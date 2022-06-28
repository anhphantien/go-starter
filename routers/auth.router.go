package routers

import (
	"go-starter/handlers"
	"go-starter/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router) {
	r.Handle(
		"/auth/login",
		middlewares.NewChain(
			middlewares.A,
			middlewares.B,
		).Then(
			handlers.AuthLogin,
		),
	).Methods(http.MethodPost)
}
