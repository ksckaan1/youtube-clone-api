package gormadp

import (
	"context"
	"youtube-clone/pkg/enum/liketype"
	"youtube-clone/pkg/repository/dto"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, um *dto.User) error
	GetUserByEmail(ctx context.Context, email string) (*dto.User, error)
	GetUserByID(ctx context.Context, id uint) (*dto.User, error)
}

type SubscribeRepositoryInterface interface {
	Subscribe(ctx context.Context, userID, channelID uint) error
	UnSubscribe(ctx context.Context, userID, channelID uint) error
	GetSubscribers(ctx context.Context, channelID uint, limit, offset int) (*dto.SubscribeList, error)
	GetSubscribeList(ctx context.Context, userID uint, limit, offset int) (*dto.SubscribeList, error)
}

type VideoRepositoryInterface interface {
	CreateVideo(ctx context.Context, vid *dto.Video) error
	GetVideo(ctx context.Context, id uint) (*dto.Video, error)
	GetAllVideos(ctx context.Context, limit, offset int) (*dto.VideoList, error)
	GetVideosByChannelID(ctx context.Context, channelID uint, limit, offset int) (*dto.VideoList, error)
	DeleteVideo(ctx context.Context, id uint) error
}

type ViewRepositoryInterface interface {
	CreateView(ctx context.Context, vw *dto.View) error
	GetViewsByVideoID(ctx context.Context, videoID uint, limit, offset int) (*dto.ViewList, error)
	GetViewsByUserID(ctx context.Context, userID uint, limit, offset int) (*dto.ViewList, error)
}

type LikeRepositoryInterface interface {
	CreateLike(ctx context.Context, lk *dto.Like) error
	DeleteLike(ctx context.Context, userID, videoID uint) error
	GetAllLikesByVideoID(ctx context.Context, videoID uint, t liketype.LikeType, limit, offset int) (*dto.LikeList, error)
}
