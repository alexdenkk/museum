package app

import (
	// users
	users_handler "akpl/museum/internal/users/gateway/http"
	users_repository "akpl/museum/internal/users/repository"
	users_service "akpl/museum/internal/users/service"
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
