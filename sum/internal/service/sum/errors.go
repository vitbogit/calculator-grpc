package sum

import "errors"

var (
	ErrInvalidInput   = errors.New("invalid input")
	ErrDivisionByZero = errors.New("division by zero")
)
