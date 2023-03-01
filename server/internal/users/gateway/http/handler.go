package http

import (
	"akpl/museum/internal/users"
)

type Handler struct {
	Service users.Service
}

func New(service users.Service) *Handler {
	return &Handler{
		Service: service,
	}
}
