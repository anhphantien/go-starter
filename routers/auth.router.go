package routers

import (
	"go-starter/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router) {
	r.HandleFunc("/auth/login", handlers.AuthHandler{}.Login).
		Methods(http.MethodPost)
}
