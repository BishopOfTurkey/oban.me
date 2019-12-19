package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Printf("%v - %vms - %v", r.URL.Path, time.Since(start).Milliseconds(), "(STATUS CODE)")
			}()

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
