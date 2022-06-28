package middlewares

import "net/http"

type middleware func(http.HandlerFunc) http.HandlerFunc

type middlewareChain []middleware

func NewChain(middlewares ...middleware) middlewareChain {
	return append(middlewareChain{}, middlewares...)
}

func (c middlewareChain) Then(h http.HandlerFunc) http.HandlerFunc {
	for i := range c {
		h = c[len(c)-1-i](h)
	}
	return h
}
