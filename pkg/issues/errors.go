package issues

import (
	"errors"
)

var (
	ErrInvalidCommand       = errors.New("not a valid command")
	ErrInvalidKey           = errors.New("key does not exist")
	ErrInvalidArgumentCount = errors.New("invalid number of arguments")
	ErrInvalidArguments     = errors.New("not valid arguments")
)
