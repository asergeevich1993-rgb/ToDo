package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("запрос %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("Время запроса %s", time.Since(start))
	})

}
