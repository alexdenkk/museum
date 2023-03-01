package service

import (
	"akpl/museum/internal/users"
	"akpl/museum/model"
	"akpl/museum/pkg/hash"
	"akpl/museum/pkg/token"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (s *Service) Get(ctx context.Context, id uint) (model.User, error) {
	user, err := s.Repository.Get(ctx, id)

	if err != nil {
		return model.User{}, err
	}

	user.Password = ""
	return user, nil
}

func (s *Service) Create(ctx context.Context, user model.User, act *token.Claims) error {
	if act.Role != model.ADMIN {
		return users.PermissionsError
	}

	if len(user.Password) < 8 || len(user.Login) < 4 {
		return users.ShortFieldError
	}

	// check if user with this login already exists
	if _, err := s.Repository.GetByLogin(ctx, user.Login); err == nil {
		return users.LoginAlreadyExistsError
	}

	// hash password
	user.Password = hash.Hash(user.Password)

	return s.Repository.Create(ctx, user)
}

func (s *Service) Update(ctx context.Context, user model.User, act *token.Claims) error {
	// check user presense
	prev, err := s.Repository.Get(ctx, user.ID)

	if err != nil {
		return err
	}

	// block changing roles for non admin users
	if act.Role != model.ADMIN && prev.Role != user.Role {
		return users.PermissionsError
	}

	// block editing admins
	if prev.Role == model.ADMIN && act.ID != user.ID {
		return users.PermissionsError
	}

	// block password changing
	if user.Password != "" && prev.Password != hash.Hash(user.Password) {
		return users.ChangePasswordError
	}

	if len(user.Login) < 4 {
		return users.ShortFieldError
	}

	return s.Repository.Update(ctx, user)
}

func (s *Service) Delete(ctx context.Context, id uint, act *token.Claims) error {
	// if user want to edit himself, he must put 0 id
	if id == 0 {
		id = act.ID
	}

	// check user presense
	user, err := s.Repository.Get(ctx, id)

	if err != nil {
		return err
	}

	// block deleting another users for non admins
	if act.Role != model.ADMIN && act.ID != id {
		return users.PermissionsError
	}

	// block deleting admins
	if user.Role == model.ADMIN {
		return users.PermissionsError
	}

	return s.Repository.Delete(ctx, id)
}

func (s *Service) Login(ctx context.Context, login, password string) (string, error) {
	// check user presense by login
	user, err := s.Repository.GetByLogin(ctx, login)

	if err != nil {
		return "", err
	}

	// check password
	if user.Password != hash.Hash(password) {
		return "", users.IncorrectPasswordError
	}

	// generating token
	claims := token.Claims{
		ID:    user.ID,
		Login: user.Login,
		Role:  user.Role,

		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1000).Unix(),
		},
	}

	return token.GenerateJWT(claims, s.SignKey)
}

func (s *Service) Register(ctx context.Context, login, password string) error {
	// check if user with this login already exists
	if _, err := s.Repository.GetByLogin(ctx, login); err == nil {
		return users.LoginAlreadyExistsError
	}

	if len(password) < 6 || len(login) < 4 {
		return users.ShortFieldError
	}

	// register new user
	user := model.User{
		Login:    login,
		Password: hash.Hash(password),
		Role:     model.USER,
	}

	return s.Repository.Create(ctx, user)
}

func (s *Service) ChangePassword(ctx context.Context, newPswd, pswd string, act *token.Claims) error {
	// check user presense
	user, err := s.Repository.Get(ctx, act.ID)

	if err != nil {
		return err
	}

	// check password
	if hash.Hash(pswd) != user.Password {
		return users.IncorrectPasswordError
	}

	if len(newPswd) < 8 {
		return users.ShortFieldError
	}

	// set new password
	user.Password = hash.Hash(newPswd)

	return s.Repository.Update(ctx, user)
}
