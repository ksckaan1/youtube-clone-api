package contextmanager

import (
	"context"
	"fmt"
)

var _ Interface = (*ContextManager)(nil)

type ContextManager struct {
	userID     uint
	isLoggedIn bool
}

func New(ctx context.Context) (*ContextManager, error) {
	var (
		userID     uint
		isLoggedIn bool
		ok         bool
	)

	if userID, ok = ctx.Value(CtxUserID).(uint); !ok {
		return nil, fmt.Errorf("new context manager / id : %w", ErrContextValueNotFound)
	}

	if isLoggedIn, ok = ctx.Value(CtxIsLoggedIn).(bool); !ok {
		return nil, fmt.Errorf("new context manager / id : %w", ErrContextValueNotFound)
	}

	return &ContextManager{
		userID:     userID,
		isLoggedIn: isLoggedIn,
	}, nil
}

func (c *ContextManager) GetUserID() uint {
	return c.userID
}

func (c *ContextManager) IsLoggedIn() bool {
	return c.isLoggedIn
}
