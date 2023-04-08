package gormadp

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"youtube-clone/pkg/repository/dto"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
)

var _ ViewRepositoryInterface = (*ViewRepository)(nil)

type ViewRepository struct {
	db *gorm.DB
}

func NewViewRepository(db *gorm.DB) *ViewRepository {
	return &ViewRepository{
		db: db,
	}
}

func (v ViewRepository) CreateView(ctx context.Context, vw *dto.View) error {
	var (
		dbView dbmodels.View
		err    error
	)

	dbView.From(vw)

	// avoid duplicate
	dbView.Key = fmt.Sprintf("%d:%d", vw.UserID, vw.VideoID)

	if err = v.db.WithContext(ctx).
		Create(&dbView).Error; err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return fmt.Errorf("create view / create : %w", ErrDuplicate)
		}

		return fmt.Errorf("create view / create : %w", err)
	}

	*vw = *dbView.To()

	return nil
}

func (v ViewRepository) GetViewsByVideoID(ctx context.Context, videoID uint, limit, offset int) (*dto.ViewList, error) {
	var (
		dbViews []dbmodels.View
		count   int64
		err     error
	)

	if err = v.db.WithContext(ctx).
		Model(&dbmodels.View{}).
		Where("video_id = ?", videoID).
		Limit(limit).Offset(offset).Count(&count).
		Find(&dbViews).Error; err != nil {
		return nil, fmt.Errorf("get views by video id / find : %w", err)
	}

	list := make([]*dto.View, 0, len(dbViews))

	for i := range dbViews {
		list = append(list, dbViews[i].To())
	}

	return &dto.ViewList{
		Limit:  limit,
		Offset: offset,
		Count:  count,
		Views:  list,
	}, nil
}

func (v ViewRepository) GetViewsByUserID(ctx context.Context, userID uint, limit, offset int) (*dto.ViewList, error) {
	var (
		dbViews []dbmodels.View
		count   int64
		err     error
	)

	if err = v.db.WithContext(ctx).
		Model(&dbmodels.View{}).
		Where("user_id = ?", userID).
		Limit(limit).Offset(offset).Count(&count).
		Find(&dbViews).Error; err != nil {
		return nil, fmt.Errorf("get views by user id / find : %w", err)
	}

	list := make([]*dto.View, 0, len(dbViews))

	for i := range dbViews {
		list = append(list, dbViews[i].To())
	}

	return &dto.ViewList{
		Limit:  limit,
		Offset: offset,
		Count:  count,
		Views:  list,
	}, nil
}
