package service

import (
	"akpl/museum/internal/places"
	"akpl/museum/model"
	"akpl/museum/pkg/token"
	"context"
)

func (s *Service) GetAll(ctx context.Context) ([]model.Place, error) {
	return s.Repository.GetAll(ctx)
}

func (s *Service) Get(ctx context.Context, id uint) (model.Place, error) {
	return s.Repository.Get(ctx, id)
}

func (s *Service) Create(ctx context.Context, place model.Place, act *token.Claims) error {
	if act.Role != model.ADMIN {
		return places.PermissionsError
	}

	if len(place.Name) < 10 || len(place.Description) < 20 || len(place.Address) < 10 {
		return places.ShortFieldError
	}

	return s.Repository.Create(ctx, place)
}

func (s *Service) Update(ctx context.Context, place model.Place, act *token.Claims) error {
	if act.Role != model.ADMIN {
		return places.PermissionsError
	}

	if len(place.Name) < 10 || len(place.Description) < 20 || len(place.Address) < 10 {
		return places.ShortFieldError
	}

	return s.Repository.Update(ctx, place)
}

func (s *Service) Delete(ctx context.Context, id uint, act *token.Claims) error {
	if act.Role != model.ADMIN {
		return places.PermissionsError
	}

	return s.Repository.Delete(ctx, id)
}
