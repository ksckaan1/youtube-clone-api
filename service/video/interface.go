package video

import "context"

type Interface interface {
	CreateVideo(ctx context.Context, vid *CreateVideoModel) (uint, error)
	GetVideo(ctx context.Context, id uint) (*GetVideoModel, error)
}
