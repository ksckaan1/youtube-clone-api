package gormadp

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"youtube-clone/pkg/repository/dto"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
)

var _ SubscribeRepositoryInterface = (*SubscribeRepository)(nil)

type SubscribeRepository struct {
	db *gorm.DB
}

func NewSubscribeRepository(db *gorm.DB) *SubscribeRepository {
	return &SubscribeRepository{
		db: db,
	}
}

func (s SubscribeRepository) Subscribe(ctx context.Context, userID, channelID uint) error {
	dbSubscribe := dbmodels.Subscribe{
		UserID:    userID,
		ChannelID: channelID,
		Key:       fmt.Sprintf("%d:%d", userID, channelID),
	}

	if err := s.db.WithContext(ctx).
		Create(&dbSubscribe).Error; err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return fmt.Errorf("subscribe : %w", ErrDuplicate)
		}

		return fmt.Errorf("subscribe : %w", err)
	}

	return nil
}

func (s SubscribeRepository) UnSubscribe(ctx context.Context, userID, channelID uint) error {
	tx := s.db.WithContext(ctx).
		Where("user_id = ? and channel_id = ?", userID, channelID).
		Delete(&dbmodels.Subscribe{})

	if tx.Error != nil {
		return fmt.Errorf("unsubscribe : %w", tx.Error)
	}

	if tx.RowsAffected == 0 {
		return fmt.Errorf("unsubscribe : %w", gorm.ErrRecordNotFound)
	}

	return nil
}

func (s SubscribeRepository) GetSubscribers(ctx context.Context, channelID uint, limit, offset int) (*dto.SubscribeList, error) {
	var (
		dbSubscribes []dbmodels.Subscribe
		count        int64
		err          error
	)

	if err = s.db.WithContext(ctx).
		Model(&dbmodels.Subscribe{}).
		Where("channel_id = ?", channelID).
		Limit(limit).Offset(offset).
		Count(&count).
		Find(&dbSubscribes).Error; err != nil {
		return nil, fmt.Errorf("get subscribers : %w", err)
	}

	list := make([]*dto.Subscribe, 0, len(dbSubscribes))

	for i := range dbSubscribes {
		list = append(list, &dto.Subscribe{
			UserID:    dbSubscribes[i].UserID,
			ChannelID: dbSubscribes[i].ChannelID,
			Key:       dbSubscribes[i].Key,
		})
	}

	return &dto.SubscribeList{
		Limit:      limit,
		Offset:     offset,
		Count:      count,
		Subscribes: list,
	}, nil
}

func (s SubscribeRepository) GetSubscribeList(ctx context.Context, userID uint, limit, offset int) (*dto.SubscribeList, error) {
	var (
		dbSubscribes []dbmodels.Subscribe
		count        int64
		err          error
	)

	if err = s.db.WithContext(ctx).
		Model(&dbmodels.Subscribe{}).
		Where("user_id = ?", userID).
		Limit(limit).Offset(offset).
		Count(&count).
		Find(&dbSubscribes).Error; err != nil {
		return nil, fmt.Errorf("get subscribe list : %w", err)
	}

	list := make([]*dto.Subscribe, 0, len(dbSubscribes))

	for i := range dbSubscribes {
		list = append(list, &dto.Subscribe{
			UserID:    dbSubscribes[i].UserID,
			ChannelID: dbSubscribes[i].ChannelID,
			Key:       dbSubscribes[i].Key,
		})
	}

	return &dto.SubscribeList{
		Limit:      limit,
		Offset:     offset,
		Count:      count,
		Subscribes: list,
	}, nil
}
