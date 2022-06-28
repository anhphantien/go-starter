package routers

import (
	"go-starter/handlers"
	"go-starter/middlewares"

	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router) {
	r.Handle("/auth/login", middlewares.
		NewChain(
			middlewares.A,
			middlewares.B,
		).
		Then(handlers.AuthLogin),
	)
}
