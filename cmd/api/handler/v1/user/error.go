package user

import "errors"

var (
	ErrServiceShouldBeSetUp = errors.New("service should be set up")
	ErrInvalidParam         = errors.New("invalid parameter")
)
