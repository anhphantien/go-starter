package routers

import (
	"go-starter/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

var authHandler = handlers.AuthHandler{}

func AuthRouter(r *mux.Router) {
	r.HandleFunc("/auth/login", authHandler.Login).
		Methods(http.MethodPost)
}
