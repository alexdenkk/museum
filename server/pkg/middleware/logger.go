package middleware

import (
	"net/http"
	"log"
	"strings"
	"time"
	"fmt"
)

func LoggerMW(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)

			log.Println(
				strings.Join([]string{
					r.Method,
					r.RemoteAddr,
					r.URL.EscapedPath(),
					fmt.Sprint(time.Since(start)),
				}, " "),
			)
		},
	)
}
