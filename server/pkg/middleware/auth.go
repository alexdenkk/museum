package middleware

import (
	"akpl/museum/pkg/token"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, *token.Claims)

func (mw *Middleware) Auth(f HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		t, err := r.Cookie("token")

		if err != nil {
			http.Redirect(w, r, "/user/register/", http.StatusMovedPermanently)
			return
		}

		claims, err := token.ParseJWT(t.Value, mw.SignKey)

		if err != nil {
			http.Redirect(w, r, "/user/register/", http.StatusMovedPermanently)
			return
		}

		f(w, r, claims)
	})
}

func (mw *Middleware) NotAuth(f http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		t, err := r.Cookie("token")

		if err != nil {
			f(w, r)
			return
		}

		_, err = token.ParseJWT(t.Value, mw.SignKey)

		if err != nil {
			f(w, r)
			return
		}

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	})
}
