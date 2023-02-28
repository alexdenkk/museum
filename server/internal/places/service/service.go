package service

import (
	"alexdenkk/books/internal/books"
)

type Service struct {
	Repository books.Repository
}

func New(repo books.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}
