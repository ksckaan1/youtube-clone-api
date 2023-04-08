package subscribe

import "context"

type Interface interface {
	Subscribe(ctx context.Context, userID, channelID uint) error
	Unsubscribe(ctx context.Context, userID, channelID uint) error
	GetSubscribers(ctx context.Context, channelID uint, limit, offset int) (*SubList, error)
	GetSubscribeList(ctx context.Context, userID uint, limit, offset int) (*SubList, error)
}
