package gormadp

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"testing"
	"youtube-clone/pkg/enum/liketype"
	"youtube-clone/pkg/repository/dto"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
)

func TestLikeRepository_CreateLike(t *testing.T) {
	db := NewTestDBConnection(t)
	likeRepo := NewLikeRepository(db)

	testLikes := []dto.Like{
		{
			UserID:  1,
			VideoID: 1,
			Type:    liketype.Like,
		},
		{
			UserID:  1,
			VideoID: 1,
			Type:    liketype.Dislike,
		},
		{
			UserID:  2,
			VideoID: 1,
			Type:    liketype.Like,
		},
		{
			UserID:  2,
			VideoID: 1,
			Type:    liketype.Dislike,
		},
	}

	for i := range testLikes {
		t.Run(fmt.Sprintf("CreateLike_Index:%d", i), func(t *testing.T) {
			if err := likeRepo.CreateLike(context.Background(), &testLikes[i]); err != nil {
				t.Error(err)
			}
		})
	}

	t.Run("CheckDuplicate", func(t *testing.T) {
		var dbLikes []dbmodels.Like

		if err := db.Find(&dbLikes).Error; err != nil {
			t.Error(err)
		}

		if len(dbLikes) != 2 {
			t.Error("like count not correct")
		}
	})
}

func TestLikeRepository_DeleteLike(t *testing.T) {
	db := NewTestDBConnection(t)
	likeRepo := NewLikeRepository(db)

	testLike := dbmodels.Like{
		UserID:  1,
		VideoID: 1,
		Type:    1,
		Key:     "1:1",
	}

	t.Run("CreateLike", func(t *testing.T) {
		if err := db.Create(&testLike).Error; err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteLike", func(t *testing.T) {
		if err := likeRepo.DeleteLike(context.Background(), testLike.UserID, testLike.VideoID); err != nil {
			t.Error(err)
		}
	})

	t.Run("CheckLikeIsDeleted", func(t *testing.T) {
		if err := db.
			Where("user_id = ? and video_id = ?", testLike.UserID, testLike.VideoID).
			First(&dbmodels.Like{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Error(err)
		}
	})

}

func TestLikeRepository_GetAllLikesByVideoID(t *testing.T) {
	db := NewTestDBConnection(t)
	likeRepo := NewLikeRepository(db)

	testLikes := []dbmodels.Like{
		{
			UserID:  1,
			VideoID: 1,
			Type:    0,
			Key:     "1:1",
		},
		{
			UserID:  2,
			VideoID: 1,
			Type:    0,
			Key:     "2:1",
		},
		{
			UserID:  3,
			VideoID: 1,
			Type:    0,
			Key:     "3:1",
		},
		{
			UserID:  4,
			VideoID: 1,
			Type:    0,
			Key:     "4:1",
		},
		{
			UserID:  5,
			VideoID: 1,
			Type:    0,
			Key:     "5:1",
		},
	}

	for i := range testLikes {
		t.Run(fmt.Sprintf("CreateLike_Index:%d", i), func(t *testing.T) {
			if err := db.Create(&testLikes[i]).Error; err != nil {
				t.Error(err)
			}
		})
	}

	t.Run("GetAllLikesByVideoID", func(t *testing.T) {
		var (
			likes *dto.LikeList
			err   error
		)

		if likes, err = likeRepo.GetAllLikesByVideoID(context.Background(), 1, liketype.Like, -1, 0); err != nil {
			t.Error(err)
		}

		if len(likes.Likes) != len(testLikes) {
			t.Error("like count not matched")
		}
	})
}