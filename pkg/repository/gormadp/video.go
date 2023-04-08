package gormadp

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"youtube-clone/pkg/repository/dto"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
)

var _ VideoRepositoryInterface = (*VideoRepository)(nil)

type VideoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepository {
	return &VideoRepository{
		db: db,
	}
}

func (v VideoRepository) CreateVideo(ctx context.Context, vid *dto.Video) error {
	var (
		dbVideo dbmodels.Video
		err     error
	)

	dbVideo.From(vid)

	if err = v.db.WithContext(ctx).
		Create(&dbVideo).Error; err != nil {
		return fmt.Errorf("create video / create : %w", err)
	}

	*vid = *dbVideo.To()

	return nil
}

func (v VideoRepository) GetVideo(ctx context.Context, id uint) (*dto.Video, error) {
	var (
		dbVideo dbmodels.Video
		err     error
	)

	if err = v.db.WithContext(ctx).
		First(&dbVideo, id).Error; err != nil {
		return nil, fmt.Errorf("get video / first : %w", err)
	}

	return dbVideo.To(), nil
}

func (v VideoRepository) GetAllVideos(ctx context.Context, limit, offset int) (*dto.VideoList, error) {
	var (
		dbVideos []dbmodels.Video
		count    int64
		err      error
	)

	if err = v.db.WithContext(ctx).
		Model(&dbmodels.Video{}).
		Count(&count).Limit(limit).Offset(offset).
		Find(&dbVideos).Error; err != nil {
		return nil, fmt.Errorf("get all videos / find : %w", err)
	}

	list := make([]*dto.Video, 0, len(dbVideos))

	for i := range dbVideos {
		list = append(list, dbVideos[i].To())
	}

	return &dto.VideoList{
		Limit:  limit,
		Offset: offset,
		Count:  count,
		Videos: list,
	}, nil
}

func (v VideoRepository) GetVideosByChannelID(ctx context.Context, channelID uint, limit, offset int) (*dto.VideoList, error) {
	var (
		dbVideos []dbmodels.Video
		count    int64
		err      error
	)

	if err = v.db.WithContext(ctx).
		Model(&dbmodels.Video{}).
		Where("user_id = ?", channelID).
		Count(&count).Limit(limit).Offset(offset).
		Find(&dbVideos).Error; err != nil {
		return nil, fmt.Errorf("get all videos by channel id / find : %w", err)
	}

	list := make([]*dto.Video, 0, len(dbVideos))

	for i := range dbVideos {
		list = append(list, dbVideos[i].To())
	}

	return &dto.VideoList{
		Limit:  limit,
		Offset: offset,
		Count:  count,
		Videos: list,
	}, nil
}

func (v VideoRepository) DeleteVideo(ctx context.Context, id uint) error {
	tx := v.db.WithContext(ctx).
		Delete(&dbmodels.Video{}, id)

	if tx.Error != nil {
		return fmt.Errorf("delete video / delete : %w", tx.Error)
	}

	if tx.RowsAffected == 0 {
		return fmt.Errorf("delete video / rows affected : %w", gorm.ErrRecordNotFound)
	}

	return nil
}
