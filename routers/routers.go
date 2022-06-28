package routers

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func New() {
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	AuthRouter()
}
