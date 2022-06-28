package routers

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func New() *mux.Router {
	r := mux.NewRouter()
	swaggerInit(r)
	apiVersion1(r)
	return r
}

func swaggerInit(r *mux.Router) {
	r.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		var scheme string
		if r.TLS == nil {
			scheme = "http"
		} else {
			scheme = "https"
		}
		http.Redirect(w, r, scheme+"://"+path.Join(r.Host, r.URL.Path, "index.html"), http.StatusMovedPermanently)
	})
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}

func apiVersion1(r *mux.Router) {
	s := r.PathPrefix("/api/v1").Subrouter()
	AuthRouter(s)
}
