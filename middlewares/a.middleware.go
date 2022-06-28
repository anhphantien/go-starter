package middlewares

import (
	"fmt"
	"net/http"
)

func A(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(111111111111)
		h.ServeHTTP(w, r)
	})
}
