package app

import (
	"akpl/museum/pkg/middleware"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *App) Route() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))))

	app.RouteUsers(r)
	app.RoutePlaces(r)
	app.RouteEvents(r)

	r.Use(middleware.LoggerMW)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp, err := template.ParseFiles("web/templates/base.html", "web/templates/pages/index.html")

		if err != nil {
			log.Fatal(err)
		}

		tmp.Execute(w, nil)
	})

	app.Server.Handler = r
}

func (app *App) RouteUsers(r *mux.Router) {
	sub := r.PathPrefix("/user").Subrouter()

	sub.HandleFunc("/register/", app.MW.NotAuth(app.UsersHandler.Register))
	sub.HandleFunc("/login/", app.MW.NotAuth(app.UsersHandler.Login))
}

func (app *App) RoutePlaces(r *mux.Router) {
	sub := r.PathPrefix("/place").Subrouter()

	sub.HandleFunc("/all/", app.PlacesHandler.GetAll)
	sub.HandleFunc("/{id:[0-9]+}/", app.PlacesHandler.Get).Methods("GET")
	sub.HandleFunc("/add/", app.MW.Auth(app.PlacesHandler.Create))
}

func (app *App) RouteEvents(r *mux.Router) {
	sub := r.PathPrefix("/event").Subrouter()

	sub.HandleFunc("/all/", app.EventsHandler.GetAll)
	sub.HandleFunc("/{id:[0-9]+}/", app.EventsHandler.Get).Methods("GET")
	sub.HandleFunc("/add/", app.MW.Auth(app.EventsHandler.Create))
	sub.HandleFunc("/search/", app.EventsHandler.Search)
	sub.HandleFunc("/for/{id:[0-9]+}/", app.EventsHandler.GetFor)
}
