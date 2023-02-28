package places

import (
	"akpl/museum/model"
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]model.Place, error)
	Get(ctx context.Context, id uint) (model.Place, error)
	Create(ctx context.Context, place model.Place) error
	Update(ctx context.Context, place model.Place) error
	Delete(ctx context.Context, id uint) error
}
