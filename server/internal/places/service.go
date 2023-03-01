package places

import (
	"akpl/museum/model"
	"akpl/museum/pkg/token"
	"context"
)

type Service interface {
	GetAll(ctx context.Context) ([]model.Place, error)
	Get(ctx context.Context, id uint) (model.Place, error)
	Create(ctx context.Context, place *model.Place, act *token.Claims) error
	Update(ctx context.Context, place *model.Place, act *token.Claims) error
	Delete(ctx context.Context, id uint, act *token.Claims) error
}
