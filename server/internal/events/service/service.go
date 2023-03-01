package service

import (
	"akpl/museum/internal/events"
)

type Service struct {
	SignKey    []byte
	Repository events.Repository
}

func New(repo events.Repository, key []byte) *Service {
	return &Service{
		Repository: repo,
		SignKey:    key,
	}
}
