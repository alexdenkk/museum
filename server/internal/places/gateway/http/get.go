package http

import (
	"akpl/museum/model"
	"context"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PlacePage struct {
	Place model.Place
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	ctx := context.WithValue(context.Background(), "request", r)

	place, err := h.Service.Get(ctx, uint(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	tmp, err := template.ParseFiles("web/templates/base.html", "web/templates/pages/place.html")

	if err != nil {
		log.Fatal(err)
	}

	tmp.Execute(w, PlacePage{Place: place})
}
