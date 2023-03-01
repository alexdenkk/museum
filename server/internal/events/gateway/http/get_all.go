package http

import (
	"akpl/museum/model"
	"context"
	"html/template"
	"log"
	"net/http"
)

type AllEventsPage struct {
	Events []model.Event
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(context.Background(), "request", r)

	events, err := h.Service.GetAll(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmp, err := template.ParseFiles("web/templates/base.html", "web/templates/pages/events.html")

	if err != nil {
		log.Fatal(err)
	}

	tmp.Execute(w, AllEventsPage{Events: events})
}
