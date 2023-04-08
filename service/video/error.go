package video

import "errors"

var (
	ErrRepositoryShouldBeSetUp = errors.New("repository should be set up")
	ErrNoPermission            = errors.New("no permission")
)
