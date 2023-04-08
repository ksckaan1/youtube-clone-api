package user

import (
	"context"
)

type Interface interface {
	CreateUser(ctx context.Context, m *CreateUserModel) error
	GetUserByID(ctx context.Context, userID uint) (*GetUserModel, error)
}
