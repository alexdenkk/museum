package repository

import (
	"akpl/museum/model"
	"context"
)

func (r *Repository) Get(ctx context.Context, id uint) (model.User, error) {
	var user model.User
	result := r.DB.First(&user, id)
	return user, result.Error
}

func (r *Repository) Create(ctx context.Context, user model.User) error {
	return r.DB.Create(&user).Error
}

func (r *Repository) Update(ctx context.Context, user model.User) error {
	return r.DB.Save(&user).Error
}

func (r *Repository) Delete(ctx context.Context, id uint) error {
	return r.DB.Delete(&model.User{}, id).Error
}

func (r *Repository) GetByLogin(ctx context.Context, login string) (model.User, error) {
	var user model.User
	result := r.DB.Where("login = ?", login).First(&user)
	return user, result.Error
}
