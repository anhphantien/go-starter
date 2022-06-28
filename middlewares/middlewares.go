package middlewares

import "net/http"

type middleware func(http.HandlerFunc) http.HandlerFunc

type chain []middleware

func NewChain(middlewares ...middleware) chain {
	return append(chain{}, middlewares...)
}

func (c chain) Then(h http.HandlerFunc) http.HandlerFunc {
	return h
}
