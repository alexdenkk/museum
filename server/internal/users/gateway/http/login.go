package http

import (
	"akpl/museum/pkg/cookie"
	"context"
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmp, err := template.ParseFiles("web/templates/base.html", "web/templates/pages/login.html")

		if err != nil {
			log.Fatal(err)
		}

		tmp.Execute(w, nil)
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		login := r.Form.Get("login")
		password := r.Form.Get("password")

		ctx := context.WithValue(context.Background(), "request", r)

		token, err := h.Service.Login(ctx, login, password)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cookie.SetCookie("token", token, w)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
