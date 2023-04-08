package video

import (
	"context"
	"errors"
	"fmt"
	"youtube-clone/pkg/contextmanager"
	"youtube-clone/pkg/enum/liketype"
	"youtube-clone/pkg/repository/dto"
	"youtube-clone/pkg/repository/gormadp"
)

var _ Interface = (*Video)(nil)

type Video struct {
	repo *gormadp.Repository
}

func New(r *gormadp.Repository) (*Video, error) {
	if r == nil {
		return nil, fmt.Errorf("new subscribe service : %w", ErrRepositoryShouldBeSetUp)
	}

	return &Video{
		repo: r,
	}, nil
}

func (v *Video) CreateVideo(ctx context.Context, vid *CreateVideoModel) (uint, error) {
	var (
		cm  *contextmanager.ContextManager
		err error
	)

	if cm, err = contextmanager.New(ctx); err != nil {
		return 0, fmt.Errorf("create video / cm : %w", err)
	}

	m := dto.Video{
		Title:       vid.Title,
		Description: vid.Description,
		URL:         vid.URL,
		Thumbnail:   vid.Thumbnail,
		UserID:      cm.GetUserID(),
	}

	if err = v.repo.VideoRepo.CreateVideo(ctx, &m); err != nil {
		return 0, fmt.Errorf("create video / create : %w", err)
	}

	return m.ID, nil
}

func (v *Video) GetVideo(ctx context.Context, id uint) (*GetVideoModel, error) {
	var (
		cm  *contextmanager.ContextManager
		err error
	)

	if cm, err = contextmanager.New(ctx); err != nil {
		return nil, fmt.Errorf("get video / cm : %w", err)
	}

	var dbVideo *dto.Video

	if dbVideo, err = v.repo.VideoRepo.GetVideo(ctx, id); err != nil {
		return nil, fmt.Errorf("get video / get : %w", err)
	}

	if cm.IsLoggedIn() {
		view := dto.View{
			UserID:  cm.GetUserID(),
			VideoID: id,
		}

		if err = v.repo.ViewRepo.CreateView(ctx, &view); err != nil {
			if !errors.Is(err, gormadp.ErrDuplicate) {
				return nil, fmt.Errorf("get video / create view : %w", err)
			}
		}
	}

	var views *dto.ViewList

	if views, err = v.repo.ViewRepo.GetViewsByVideoID(ctx, id, -1, 0); err != nil {
		return nil, fmt.Errorf("get video / get views : %w", err)
	}

	var likes *dto.LikeList

	if likes, err = v.repo.LikeRepo.GetAllLikesByVideoID(ctx, id, liketype.Like, -1, 0); err != nil {
		return nil, fmt.Errorf("get video / get likes : %w", err)
	}

	var dislikes *dto.LikeList

	if dislikes, err = v.repo.LikeRepo.GetAllLikesByVideoID(ctx, id, liketype.Dislike, -1, 0); err != nil {
		return nil, fmt.Errorf("get video / get dislikes : %w", err)
	}

	var (
		isLiked    bool
		isDisliked bool
	)

	if cm.IsLoggedIn() {
		for i := range likes.Likes {
			if likes.Likes[i].UserID == cm.GetUserID() {
				isLiked = true

				break
			}
		}

		for i := range dislikes.Likes {
			if dislikes.Likes[i].UserID == cm.GetUserID() {
				isDisliked = true

				break
			}
		}
	}

	return &GetVideoModel{
		Id:          dbVideo.ID,
		CreatedAt:   dbVideo.CreatedAt,
		Title:       dbVideo.Title,
		Description: dbVideo.Description,
		ViewCount:   views.Count,
		LikeCount:   int64(len(likes.Likes)),
		IsLiked:     isLiked,
		IsDisliked:  isDisliked,
		Comments:    nil,
	}, nil
}
