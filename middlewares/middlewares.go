package middlewares

import (
	"net/http"

	"github.com/samber/lo"
)

type middleware func(http.Handler) http.Handler

type middlewareChain []middleware

func NewChain(middlewares ...middleware) middlewareChain {
	return lo.Reverse(middlewares)
}

func (c middlewareChain) Then(h http.HandlerFunc) http.HandlerFunc {
	for _, m := range c {
		if m == nil {
			return h
		}
		h = m(h).ServeHTTP
	}
	return h
}
