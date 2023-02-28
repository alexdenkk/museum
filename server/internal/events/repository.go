package events

import (
	"akpl/museum/model"
	"context"
)

type Repository interface {
	Search(ctx context.Context, query string) ([]model.Event, error)
	GetAll(ctx context.Context) ([]model.Event, error)
	Get(ctx context.Context, id uint) (model.Event, error)
	Create(ctx context.Context, event model.Event) error
	Update(ctx context.Context, event model.Event) error
	Delete(ctx context.Context, id uint) error
}
