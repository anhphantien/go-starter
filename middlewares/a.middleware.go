package middlewares

import (
	"fmt"
	"net/http"
)

func A(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(111111111111)
		h.ServeHTTP(w, r)
	})
}
