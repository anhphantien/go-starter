package middlewares

import (
	"fmt"
	"net/http"
)

func B(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(222222222222)
		h.ServeHTTP(w, r)
	})
}
