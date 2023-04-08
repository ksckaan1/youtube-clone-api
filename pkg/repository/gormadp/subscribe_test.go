package gormadp

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"testing"
	"time"
	"youtube-clone/pkg/repository/dto"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
)

func TestSubscribeRepository_Subscribe(t *testing.T) {
	db := NewTestDBConnection(t)
	subRepo := NewSubscribeRepository(db)

	tests := []struct {
		description   string
		userId        uint
		channelID     uint
		expectedError error
	}{
		{
			description:   "Valid1",
			userId:        1,
			channelID:     1,
			expectedError: nil,
		},
		{
			description:   "Valid2",
			userId:        1,
			channelID:     2,
			expectedError: nil,
		},
		{
			description:   "Duplicate",
			userId:        1,
			channelID:     1,
			expectedError: ErrDuplicate,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Subscribe_%s", tt.description), func(t *testing.T) {
			if err := subRepo.Subscribe(context.Background(), tt.userId, tt.channelID); !errors.Is(err, tt.expectedError) {
				t.Error(err)
			}
		})
	}

	t.Run("ErrorTesting", func(t *testing.T) {
		t.Run("ContextTimeout", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*0)
			defer cancel()

			if err := subRepo.Subscribe(ctx, 1, 1); err == nil {
				t.Error("expected: context deadline exceed")
			}
		})

		t.Run("ContextCancel", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			cancel()

			if err := subRepo.Subscribe(ctx, 1, 1); err == nil {
				t.Error("expected: context canceled")
			}
		})
	})
}

func TestSubscribeRepository_UnSubscribe(t *testing.T) {
	db := NewTestDBConnection(t)
	subRepo := NewSubscribeRepository(db)

	testSubscribe := dbmodels.Subscribe{
		UserID:    1,
		ChannelID: 1,
		Key:       "1:1",
	}

	t.Run("CreateSubscribe", func(t *testing.T) {
		if err := db.Create(&testSubscribe).Error; err != nil {
			t.Error(err)
		}
	})

	tests := []struct {
		description   string
		userID        uint
		channelID     uint
		expectedError error
	}{
		{
			description:   "Valid",
			userID:        1,
			channelID:     1,
			expectedError: nil,
		},
		{
			description:   "NotExisting",
			userID:        1,
			channelID:     1,
			expectedError: gorm.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Unsubscribe_%s", tt.description), func(t *testing.T) {
			if err := subRepo.UnSubscribe(context.Background(), tt.userID, tt.channelID); !errors.Is(err, tt.expectedError) {
				t.Error(err)
			}
		})
	}

	t.Run("ErrorTesting", func(t *testing.T) {
		t.Run("ContextTimeout", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*0)
			defer cancel()

			if err := subRepo.UnSubscribe(ctx, 1, 1); err == nil {
				t.Error("expected: context deadline exceed")
			}
		})

		t.Run("ContextCancel", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			cancel()

			if err := subRepo.UnSubscribe(ctx, 1, 1); err == nil {
				t.Error("expected: context canceled")
			}
		})
	})
}

func TestSubscribeRepository_GetSubscribers(t *testing.T) {
	db := NewTestDBConnection(t)
	subRepo := NewSubscribeRepository(db)

	testSubscribes := []dbmodels.Subscribe{
		{
			UserID:    2,
			ChannelID: 1,
			Key:       "2:1",
		},
		{
			UserID:    3,
			ChannelID: 1,
			Key:       "3:1",
		},
		{
			UserID:    4,
			ChannelID: 1,
			Key:       "4:1",
		},
		{
			UserID:    5,
			ChannelID: 1,
			Key:       "5:1",
		},
		{
			UserID:    6,
			ChannelID: 1,
			Key:       "6:1",
		},
	}

	for i := range testSubscribes {
		t.Run(fmt.Sprintf("Subscribe_Index:%d", i), func(t *testing.T) {
			if err := db.Create(&testSubscribes[i]).Error; err != nil {
				t.Error(err)
			}
		})
	}

	t.Run("GetSubscribers", func(t *testing.T) {
		var (
			subs *dto.SubscribeList
			err  error
		)

		if subs, err = subRepo.GetSubscribers(context.Background(), 1, -1, 0); err != nil {
			t.Error(err)
		}

		if len(subs.Subscribes) != len(testSubscribes) {
			t.Error("subscriber count not matching")
		}
	})

	t.Run("ErrorTesting", func(t *testing.T) {
		t.Run("ContextTimeout", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*0)
			defer cancel()

			if _, err := subRepo.GetSubscribers(ctx, 1, -1, 0); err == nil {
				t.Error("expected: context deadline exceed")
			}
		})

		t.Run("ContextCancel", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			cancel()

			if _, err := subRepo.GetSubscribers(ctx, 1, -1, 0); err == nil {
				t.Error("expected: context canceled")
			}
		})
	})
}

func TestSubscribeRepository_GetSubscribeList(t *testing.T) {
	db := NewTestDBConnection(t)
	subRepo := NewSubscribeRepository(db)

	testSubscribes := []dbmodels.Subscribe{
		{
			UserID:    1,
			ChannelID: 2,
			Key:       "1:2",
		},
		{
			UserID:    1,
			ChannelID: 3,
			Key:       "1:3",
		},
		{
			UserID:    1,
			ChannelID: 4,
			Key:       "1:4",
		},
		{
			UserID:    1,
			ChannelID: 5,
			Key:       "1:5",
		},
		{
			UserID:    1,
			ChannelID: 6,
			Key:       "1:6",
		},
	}

	for i := range testSubscribes {
		t.Run(fmt.Sprintf("CreateSubscribe_%d", i), func(t *testing.T) {
			if err := db.Create(&testSubscribes[i]).Error; err != nil {
				t.Error(err)
			}
		})
	}

	t.Run("GetSubscribeList", func(t *testing.T) {
		var (
			subs *dto.SubscribeList
			err  error
		)

		if subs, err = subRepo.GetSubscribeList(context.Background(), 1, -1, 0); err != nil {
			t.Error(err)
		}

		if len(subs.Subscribes) != len(testSubscribes) {
			t.Error(err)
		}
	})

	t.Run("ErrorTesting", func(t *testing.T) {
		t.Run("ContextTimeout", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*0)
			defer cancel()

			if _, err := subRepo.GetSubscribeList(ctx, 1, -1, 0); err == nil {
				t.Error("expected: context deadline exceed")
			}
		})

		t.Run("ContextCancel", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			cancel()

			if _, err := subRepo.GetSubscribeList(ctx, 1, -1, 0); err == nil {
				t.Error("expected: context canceled")
			}
		})
	})
}
