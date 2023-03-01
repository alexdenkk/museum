package repository

import (
	"akpl/museum/model"
	"context"
)

func (r *Repository) Search(ctx context.Context, query string) ([]model.Event, error) {
	var events []model.Event
	result := r.DB.Where("name LIKE %?%", query).Find(&events)
	return events, result.Error
}

func (r *Repository) GetAll(ctx context.Context) ([]model.Event, error) {
	var events []model.Event
	result := r.DB.Find(&events)
	return events, result.Error
}

func (r *Repository) Get(ctx context.Context, id uint) (model.Event, error) {
	var event model.Event
	result := r.DB.First(&event, id)
	return event, result.Error
}

func (r *Repository) Create(ctx context.Context, event model.Event) error {
	result := r.DB.Create(&event)
	return result.Error
}

func (r *Repository) Update(ctx context.Context, event model.Event) error {
	result := r.DB.Save(&event)
	return result.Error
}

func (r *Repository) Delete(ctx context.Context, id uint) error {
	result := r.DB.Delete(&model.Event{}, id)
	return result.Error
}
