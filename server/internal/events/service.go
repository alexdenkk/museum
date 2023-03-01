package events

import (
	"akpl/museum/model"
	"akpl/museum/pkg/token"
	"context"
)

type Service interface {
	Search(ctx context.Context, query string) ([]model.Event, error)
	GetAll(ctx context.Context) ([]model.Event, error)
	Get(ctx context.Context, id uint) (model.Event, error)
	Create(ctx context.Context, event model.Event, act *token.Claims) error
	Update(ctx context.Context, event model.Event, act *token.Claims) error
	Delete(ctx context.Context, id uint) error
}
