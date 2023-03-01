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

type ForEventsPage struct {
	Events []model.Event
}

func (h *Handler) GetFor(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	ctx := context.WithValue(context.Background(), "request", r)

	events, err := h.Service.GetFor(ctx, uint(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmp, err := template.ParseFiles("web/templates/base.html", "web/templates/pages/events.html")

	if err != nil {
		log.Fatal(err)
	}

	tmp.Execute(w, ForEventsPage{Events: events})
}
