package service

import (
	"akpl/museum/internal/events"
)

type Service struct {
	Repository events.Repository
}

func New(repo events.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}
