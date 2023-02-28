package users

import (
	"errors"
)

var (
	ShortFieldError         = errors.New("слишком мало символов")
	PermissionsError        = errors.New("нет прав на выполнение операции")
	LoginAlreadyExistsError = errors.New("пользователь с этим логином уже существует")
	ChangePasswordError     = errors.New("с помощью этой функции нельзя сменить пароль")
	IncorrectPasswordError  = errors.New("неверный пароль")
)
