package app

import (
	"akpl/museum/pkg/middleware"
	"net/http"
	"time"

	"gorm.io/gorm"

	// handlers
	// events_handler "akpl/museum/internal/events/gateway/http"
	// places_handler "akpl/museum/internal/places/gateway/http"
	users_handler "akpl/museum/internal/users/gateway/http"
)

type App struct {
	UsersHandler *users_handler.Handler
	//EventsHandler *events_handler.Handler
	//PlacesHandler *places_handler.Handler
	Server  *http.Server
	MW      *middleware.Middleware
	SignKey []byte
	DB      *gorm.DB
}

func New(db *gorm.DB, key []byte, addr string) *App {
	app := &App{
		DB:      db,
		SignKey: key,
	}

	app.Server = &http.Server{
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	app.MW = middleware.New(app.SignKey)

	app.InitUsersService()

	return app
}

func (app *App) Run() error {
	app.Route()

	return app.Server.ListenAndServe()
}
