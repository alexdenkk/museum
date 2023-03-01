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
			"web/templates/pages/create_place.html",
		)

		if err != nil {
			log.Fatal(err)
		}

		tmp.Execute(w, nil)
	}

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(10 << 20)

		name := r.FormValue("name")
		address := r.FormValue("address")
		description := r.FormValue("description")

		place := &model.Place{
			Name:        name,
			Address:     address,
			Description: description,
		}

		ctx := context.WithValue(context.Background(), "request", r)

		err := h.Service.Create(ctx, place, act)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, _, err := r.FormFile("img")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tempFile, err := os.Create("web/static/preview-img/places/img-" + strconv.Itoa(int(place.ID)) + ".jpg")
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

		place.ImgAddr = "img-" + strconv.Itoa(int(place.ID)) + ".jpg"
		h.Service.Update(ctx, place, act)

		http.Redirect(w, r, "/place/"+strconv.Itoa(int(place.ID))+"/", http.StatusMovedPermanently)
	}
}
