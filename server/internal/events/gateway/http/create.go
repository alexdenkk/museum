package http

import (
	"akpl/museum/model"
	"akpl/museum/pkg/token"
	"context"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	if r.Method == http.MethodGet {
		tmp, err := template.ParseFiles(
			"web/templates/base.html",
			"web/templates/pages/create_event.html",
		)

		if err != nil {
			log.Fatal(err)
		}

		tmp.Execute(w, nil)
	}

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(10 << 20)

		place_id, _ := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		description := r.FormValue("description")

		event := &model.Event{
			PlaceID:     uint(place_id),
			Name:        name,
			Description: description,
		}

		ctx := context.WithValue(context.Background(), "request", r)

		err := h.Service.Create(ctx, event, act)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, _, err := r.FormFile("img")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tempFile, err := os.Create("web/static/preview-img/events/img-" + strconv.Itoa(int(event.ID)) + ".jpg")
		if err != nil {
			log.Fatal(err)
		}

		fileBytes, err := ioutil.ReadAll(file)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tempFile.Write(fileBytes)

		file.Close()
		tempFile.Close()

		event.ImgAddr = "img-" + strconv.Itoa(int(event.ID)) + ".jpg"
		h.Service.Update(ctx, event, act)

		http.Redirect(w, r, "/event/"+strconv.Itoa(int(event.ID))+"/", http.StatusMovedPermanently)
	}
}
