package gormadp

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"youtube-clone/pkg/enum/liketype"
	"youtube-clone/pkg/repository/dto"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
)

var _ LikeRepositoryInterface = (*LikeRepository)(nil)

type LikeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *LikeRepository {
	return &LikeRepository{
		db: db,
	}
}

func (l LikeRepository) CreateLike(ctx context.Context, lk *dto.Like) error {
	var (
		dbLike dbmodels.Like
		err    error
	)

	dbLike.From(lk)

	dbLike.Key = fmt.Sprintf("%d:%d", lk.UserID, lk.VideoID)

	if err = l.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"user_id", "video_id", "type"}), // column needed to be updated
	}).Create(&dbLike).Error; err != nil {
		return fmt.Errorf("create like / create : %w", err)
	}

	return nil
}

func (l LikeRepository) DeleteLike(ctx context.Context, userID, videoID uint) error {
	if err := l.db.WithContext(ctx).
		Where("user_id = ? and video_id = ?", userID, videoID).
		Delete(&dbmodels.Like{}).Error; err != nil {
		return fmt.Errorf("delete like : %w", err)
	}

	return nil
}

func (l LikeRepository) GetAllLikesByVideoID(ctx context.Context, videoID uint, t liketype.LikeType, limit, offset int) (*dto.LikeList, error) {
	var (
		dbLikes []dbmodels.Like
		count   int64
		err     error
	)

	if err = l.db.WithContext(ctx).
		Model(&dbmodels.Like{}).
		Limit(limit).Offset(offset).Count(&count).
		Where("video_id = ? and type = ?", videoID, t).
		Find(&dbLikes).Error; err != nil {
		return nil, fmt.Errorf("get all likes by video id / find : %w", err)
	}

	list := make([]*dto.Like, 0, len(dbLikes))

	for i := range dbLikes {
		list = append(list, dbLikes[i].To())
	}

	return &dto.LikeList{
		Limit:  limit,
		Offset: offset,
		Count:  count,
		Likes:  list,
	}, nil
}
