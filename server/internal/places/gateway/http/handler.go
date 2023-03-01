package http

import (
	"akpl/museum/internal/places"
)

type Handler struct {
	Service places.Service
}

func New(service places.Service) *Handler {
	return &Handler{
		Service: service,
	}
}
