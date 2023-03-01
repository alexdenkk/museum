package http

import (
	"context"
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmp, err := template.ParseFiles("web/templates/base.html", "web/templates/pages/register.html")

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

		err := h.Service.Register(ctx, login, password)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/user/login/", http.StatusMovedPermanently)
	}
}
