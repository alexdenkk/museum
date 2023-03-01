package app

import (
	// users
	users_handler "akpl/museum/internal/users/gateway/http"
	users_repository "akpl/museum/internal/users/repository"
	users_service "akpl/museum/internal/users/service"

	// places
	places_handler "akpl/museum/internal/places/gateway/http"
	places_repository "akpl/museum/internal/places/repository"
	places_service "akpl/museum/internal/places/service"

	// events
	events_handler "akpl/museum/internal/events/gateway/http"
	events_repository "akpl/museum/internal/events/repository"
	events_service "akpl/museum/internal/events/service"
)

// users
func (app *App) InitUsersService() {
	// repository
	repository := users_repository.New(app.DB)
	// service
	service := users_service.New(repository, app.SignKey)
	// handler
	handler := users_handler.New(service)
	app.UsersHandler = handler
}

// places
func (app *App) InitPlacesService() {
	// repository
	repository := places_repository.New(app.DB)
	// service
	service := places_service.New(repository)
	// handler
	handler := places_handler.New(service)
	app.PlacesHandler = handler
}

// events
func (app *App) InitEventsService() {
	// repository
	repository := events_repository.New(app.DB)
	// service
	service := events_service.New(repository)
	// handler
	handler := events_handler.New(service)
	app.EventsHandler = handler
}
