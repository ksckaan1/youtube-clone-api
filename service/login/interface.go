package login

import "context"

type Interface interface {
	Login(ctx context.Context, email, password string) (string, error)
}
