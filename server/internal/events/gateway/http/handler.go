package http

import (
	"akpl/museum/internal/events"
)

type Handler struct {
	Service events.Service
}

func New(service events.Service) *Handler {
	return &Handler{
		Service: service,
	}
}
