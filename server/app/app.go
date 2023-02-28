package app

import (
	"net/http"
)

type App struct {
	Server http.Server
}

func New() *App {
	return &App{}
}
