package service

import (
	"akpl/museum/internal/events"
	"akpl/museum/model"
	"akpl/museum/pkg/token"
	"context"
)

func (s *Service) Search(ctx context.Context, query string) ([]model.Event, error) {
	return s.Repository.Search(ctx, query)
}

func (s *Service) GetFor(ctx context.Context, place_id uint) ([]model.Event, error) {
	return s.Repository.GetFor(ctx, place_id)
}

func (s *Service) GetAll(ctx context.Context) ([]model.Event, error) {
	return s.Repository.GetAll(ctx)
}

func (s *Service) Get(ctx context.Context, id uint) (model.Event, error) {
	return s.Repository.Get(ctx, id)
}

func (s *Service) Create(ctx context.Context, event *model.Event, act *token.Claims) error {
	if act.Role != model.ADMIN {
		return events.PermissionsError
	}

	if len(event.Name) < 10 || len(event.Description) < 20 {
		return events.ShortFieldError
	}

	return s.Repository.Create(ctx, event)
}

func (s *Service) Update(ctx context.Context, event *model.Event, act *token.Claims) error {
	if act.Role != model.ADMIN {
		return events.PermissionsError
	}

	if len(event.Name) < 10 || len(event.Description) < 20 {
		return events.ShortFieldError
	}

	return s.Repository.Update(ctx, event)
}

func (s *Service) Delete(ctx context.Context, id uint, act *token.Claims) error {
	if act.Role != model.ADMIN {
		return events.PermissionsError
	}

	return s.Repository.Delete(ctx, id)
}
