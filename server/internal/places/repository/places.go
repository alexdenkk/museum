package repository

import (
	"akpl/museum/model"
	"context"
)

func (r *Repository) GetAll(ctx context.Context) ([]model.Place, error) {
	var places []model.Place
	result := r.DB.Find(&places)
	return places, result.Error
}

func (r *Repository) Get(ctx context.Context, id uint) (model.Place, error) {
	var place model.Place
	result := r.DB.First(&place, id)
	return place, result.Error
}

func (r *Repository) Create(ctx context.Context, place *model.Place) error {
	result := r.DB.Create(place)
	return result.Error
}

func (r *Repository) Update(ctx context.Context, place *model.Place) error {
	result := r.DB.Save(&place)
	return result.Error
}

func (r *Repository) Delete(ctx context.Context, id uint) error {
	result := r.DB.Delete(&model.Place{}, id)
	return result.Error
}
