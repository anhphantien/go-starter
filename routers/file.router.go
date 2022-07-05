package routers

import (
	"go-starter/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

var fileHandler = handlers.FileHandler{}

func FileRouter(r *mux.Router) {
	r.HandleFunc("/file/upload", fileHandler.Upload).
		Methods(http.MethodPost)
}
