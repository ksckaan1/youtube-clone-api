package video

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"testing"
	"youtube-clone/pkg/contextmanager"
	"youtube-clone/pkg/enum/liketype"
	"youtube-clone/pkg/repository/gormadp"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
)

func TestVideo_CreateVideo(t *testing.T) {
	db := gormadp.NewTestDBConnection(t)
	repo := gormadp.NewRepository(db)
	var (
		videoService *Video
		err          error
	)

	if videoService, err = New(repo); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		description   string
		userID        uint
		expectedError error
	}{
		{
			description:   "Valid",
			userID:        1,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("CreateVideo_%s", tt.description), func(t *testing.T) {
			ctx1 := context.WithValue(context.Background(), contextmanager.CtxUserID, tt.userID)
			ctx2 := context.WithValue(ctx1, contextmanager.CtxIsLoggedIn, true)

			if _, err := videoService.CreateVideo(ctx2, &CreateVideoModel{
				Title:       "test",
				URL:         "asd",
				Thumbnail:   "asda",
				Description: "sdasd",
			}); !errors.Is(err, tt.expectedError) {
				t.Error(err)
			}
		})
	}
}

func TestVideo_GetVideo(t *testing.T) {
	db := gormadp.NewTestDBConnection(t)
	repo := gormadp.NewRepository(db)
	var (
		videoService *Video
		err          error
	)

	if videoService, err = New(repo); err != nil {
		t.Fatal(err)
	}

	testVideo := dbmodels.Video{
		Title:       "video1",
		Description: "desc1",
		URL:         "url1",
		Thumbnail:   "thumbnail1",
		UserID:      1,
	}

	t.Run("CreateVideo", func(t *testing.T) {
		if err := db.Create(&testVideo).Error; err != nil {
			t.Error(err)
		}
	})

	testLikes := []dbmodels.Like{
		{
			UserID:  1,
			VideoID: testVideo.ID,
			Type:    uint(liketype.Like),
			Key:     "1:1",
		},
		{
			UserID:  2,
			VideoID: testVideo.ID,
			Type:    uint(liketype.Like),
			Key:     "2:1",
		},
		{
			UserID:  3,
			VideoID: testVideo.ID,
			Type:    uint(liketype.Like),
			Key:     "3:1",
		},
		{
			UserID:  4,
			VideoID: testVideo.ID,
			Type:    uint(liketype.Like),
			Key:     "4:1",
		},
		{
			UserID:  5,
			VideoID: testVideo.ID,
			Type:    uint(liketype.Dislike),
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

	tests := []struct {
		description    string
		userID         uint
		isLoggedIn     bool
		videoID        uint
		wantIsLiked    bool
		wantIsDisliked bool
		wantLikeCount  uint
		wantViewCount  int64
		expectedError  error
	}{
		// invalid
		{
			description:    "NotExisting",
			userID:         1,
			isLoggedIn:     true,
			videoID:        555,
			wantIsLiked:    false,
			wantIsDisliked: false,
			wantLikeCount:  0,
			wantViewCount:  0,
			expectedError:  gorm.ErrRecordNotFound,
		},
		// valid
		{
			description:    "Valid1",
			userID:         1,
			isLoggedIn:     true,
			videoID:        testVideo.ID,
			wantIsLiked:    true,
			wantIsDisliked: false,
			wantLikeCount:  4,
			wantViewCount:  1,
			expectedError:  nil,
		},
		{
			description:    "Valid2",
			userID:         5,
			isLoggedIn:     true,
			videoID:        testVideo.ID,
			wantIsLiked:    false,
			wantIsDisliked: true,
			wantLikeCount:  4,
			wantViewCount:  2,
			expectedError:  nil,
		},
		{
			description:    "Valid2",
			userID:         555,
			isLoggedIn:     false,
			videoID:        testVideo.ID,
			wantIsLiked:    false,
			wantIsDisliked: false,
			wantLikeCount:  4,
			wantViewCount:  2,
			expectedError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("GetVideo_%s", tt.description), func(t *testing.T) {
			ctx1 := context.WithValue(context.Background(), contextmanager.CtxIsLoggedIn, tt.isLoggedIn)
			ctx2 := context.WithValue(ctx1, contextmanager.CtxUserID, tt.userID)

			var (
				resp *GetVideoModel
				err  error
			)

			if resp, err = videoService.GetVideo(ctx2, tt.videoID); !errors.Is(err, tt.expectedError) {
				t.Error(err)
			}

			if err == nil {
				if resp.IsLiked != tt.wantIsLiked {
					t.Error("isLiked not matched")
				}

				if resp.IsDisliked != tt.wantIsDisliked {
					t.Error("isDisliked not matched")
				}

				if int(resp.LikeCount) != 4 {
					t.Error("likeCount not matched")
				}

				if tt.wantViewCount != resp.ViewCount {
					t.Error("viewCount not matched")
				}
			}
		})
	}
}
