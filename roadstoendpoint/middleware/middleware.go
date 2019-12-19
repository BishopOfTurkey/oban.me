package middleware

import "net/http"

type Middlewares []func(http.Handler) http.Handler

func Chain(r http.Handler, middlewares Middlewares) http.Handler {
	for _, m := range middlewares {
		r = m(r)
	}
	return r
}
