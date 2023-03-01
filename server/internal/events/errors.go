package events

import (
	"errors"
)

var (
	ShortFieldError  = errors.New("слишком мало символов")
	PermissionsError = errors.New("нет прав на выполнение операции")
)
