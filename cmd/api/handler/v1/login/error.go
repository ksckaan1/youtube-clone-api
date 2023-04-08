package login

import "errors"

var (
	ErrServiceShouldBeSetUp = errors.New("service should be set up")
	ErrWrongEmailOrPassword = errors.New("wrong email or password")
)
