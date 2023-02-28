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
