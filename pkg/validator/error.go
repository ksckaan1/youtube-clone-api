package validator

import "errors"

var (
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidRange = errors.New("invalid range")
)
