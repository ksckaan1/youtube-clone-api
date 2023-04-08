package gormadp

import (
	"context"
	"fmt"
	"testing"
	"youtube-clone/pkg/repository/dto"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
)

func TestVideoRepository_CreateVideo(t *testing.T) {
	db := NewTestDBConnection(t)
	videoRepo := NewVideoRepository(db)

	testVideos := []dto.Video{
		{
			Title:       "video1",
			Description: "desc1",
			URL:         "url1",
			Thumbnail:   "thumb1",
			UserID:      1,
		},
		{
			Title:       "video2",
			Description: "desc2",
			URL:         "url2",
			Thumbnail:   "thumb2",
			UserID:      1,
		},
		{
			Title:       "video3",
			Description: "desc3",
			URL:         "url3",
			Thumbnail:   "thumb3",
			UserID:      1,
		},
	}

	for i := range testVideos {
		t.Run(fmt.Sprintf("CreateVideo_Index:%d", i), func(t *testing.T) {
			if err := videoRepo.CreateVideo(context.Background(), &testVideos[i]); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestVideoRepository_GetVideo(t *testing.T) {
	db := NewTestDBConnection(t)
	videoRepo := NewVideoRepository(db)

	testVideo := dbmodels.Video{
		Title:       "video1",
		Description: "desc1",
		URL:         "url1",
		Thumbnail:   "thumb1",
		UserID:      1,
	}

	t.Run("CreateVideo", func(t *testing.T) {
		if err := db.Create(&testVideo).Error; err != nil {
			t.Error(err)
		}
	})

	t.Run("GetVideo", func(t *testing.T) {
		var err error

		if _, err = videoRepo.GetVideo(context.Background(), testVideo.ID); err != nil {
			t.Error(err)
		}
	})
}

func TestVideoRepository_GetAllVideos(t *testing.T) {
	db := NewTestDBConnection(t)
	videoRepo := NewVideoRepository(db)

	testVideos := []dbmodels.Video{
		{
			Title:       "video1",
			Description: "desc1",
			URL:         "url1",
			Thumbnail:   "thumb1",
			UserID:      1,
		},
		{
			Title:       "video2",
			Description: "desc2",
			URL:         "url2",
			Thumbnail:   "thumb2",
			UserID:      1,
		},
		{
			Title:       "video3",
			Description: "desc3",
			URL:         "url3",
			Thumbnail:   "thumb3",
			UserID:      1,
		},
	}

	for i := range testVideos {
		t.Run(fmt.Sprintf("CreateVideo_Index:%d", i), func(t *testing.T) {
			if err := db.Create(&testVideos[i]).Error; err != nil {
				t.Error(err)
			}
		})
	}

	t.Run("GetAllVideos", func(t *testing.T) {
		var (
			videos *dto.VideoList
			err    error
		)

		if videos, err = videoRepo.GetAllVideos(context.Background(), -1, 0); err != nil {
			t.Error(err)
		}

		if len(videos.Videos) != len(testVideos) {
			t.Error("video lengths are not matched")
		}
	})
}

func TestVideoRepository_GetVideosByChannelID(t *testing.T) {
	db := NewTestDBConnection(t)
	videoRepo := NewVideoRepository(db)

	testVideos := []dbmodels.Video{
		{
			Title:       "video1",
			Description: "desc1",
			URL:         "url1",
			Thumbnail:   "thumb1",
			UserID:      1,
		},
		{
			Title:       "video2",
			Description: "desc2",
			URL:         "url2",
			Thumbnail:   "thumb2",
			UserID:      1,
		},
		{
			Title:       "video3",
			Description: "desc3",
			URL:         "url3",
			Thumbnail:   "thumb3",
			UserID:      1,
		},
	}

	for i := range testVideos {
		t.Run(fmt.Sprintf("CreateVideo_Index:%d", i), func(t *testing.T) {
			if err := db.Create(&testVideos[i]).Error; err != nil {
				t.Error(err)
			}
		})
	}

	t.Run("GetVideosByChannelID", func(t *testing.T) {
		var (
			videos *dto.VideoList
			err    error
		)

		if videos, err = videoRepo.GetVideosByChannelID(context.Background(), 1, -1, 0); err != nil {
			t.Error(err)
		}

		if len(videos.Videos) != len(testVideos) {
			t.Error("video lengths are not matched")
		}
	})
}

func TestVideoRepository_DeleteVideo(t *testing.T) {
	db := NewTestDBConnection(t)
	videoRepo := NewVideoRepository(db)

	testVideo := dbmodels.Video{
		Title:       "video1",
		Description: "desc1",
		URL:         "url1",
		Thumbnail:   "thumb1",
		UserID:      1,
	}

	t.Run("CreateVideo", func(t *testing.T) {
		if err := db.Create(&testVideo).Error; err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteVideo", func(t *testing.T) {
		if err := videoRepo.DeleteVideo(context.Background(), testVideo.ID); err != nil {
			t.Error(err)
		}
	})
}