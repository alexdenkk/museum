package service

import (
	"akpl/museum/internal/places"
)

type Service struct {
	Repository places.Repository
}

func New(repo places.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}
