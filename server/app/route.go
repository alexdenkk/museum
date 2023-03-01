package app

import (
	"akpl/museum/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *App) Route() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))))

	app.RouteUsers(r)

	r.Use(middleware.LoggerMW)

	app.Server.Handler = r
}

func (app *App) RouteUsers(r *mux.Router) {
	sub := r.PathPrefix("/user").Subrouter()

	sub.HandleFunc("/register/", app.MW.NotAuth(app.UsersHandler.Register))
	sub.HandleFunc("/login/", app.MW.NotAuth(app.UsersHandler.Login))
}
