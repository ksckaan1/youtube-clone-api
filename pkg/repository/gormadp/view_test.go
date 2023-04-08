package gormadp

import (
	"context"
	"fmt"
	"testing"
	"youtube-clone/pkg/repository/dto"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
)

func TestViewRepository_CreateView(t *testing.T) {
	db := NewTestDBConnection(t)
	viewRepo := NewViewRepository(db)

	testViews := []dto.View{
		{
			UserID:  1,
			VideoID: 1,
		},
		{
			UserID:  1,
			VideoID: 2,
		},
		{
			UserID:  1,
			VideoID: 3,
		},
	}

	for i := range testViews {
		t.Run(fmt.Sprintf("CreateView_Index:%d", i), func(t *testing.T) {
			if err := viewRepo.CreateView(context.Background(), &testViews[i]); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestViewRepository_GetViewsByUserID(t *testing.T) {
	db := NewTestDBConnection(t)
	viewRepo := NewViewRepository(db)

	testViews := []dbmodels.View{
		{
			UserID:  1,
			VideoID: 1,
			Key:     "1:1",
		},
		{
			UserID:  1,
			VideoID: 2,
			Key:     "1:2",
		},
		{
			UserID:  1,
			VideoID: 3,
			Key:     "1:3",
		},
	}

	for i := range testViews {
		t.Run(fmt.Sprintf("CreateView_Index:%d", i), func(t *testing.T) {
			if err := db.Create(&testViews[i]).Error; err != nil {
				t.Error(err)
			}
		})
	}

	t.Run("GetViewsByUserID", func(t *testing.T) {
		var (
			views *dto.ViewList
			err   error
		)

		if views, err = viewRepo.GetViewsByUserID(context.Background(), 1, -1, 0); err != nil {
			t.Error(err)
		}

		if len(views.Views) != len(testViews) {
			t.Error("view lengths are not matched")
		}
	})
}

func TestViewRepository_GetViewsByVideoID(t *testing.T) {
	db := NewTestDBConnection(t)
	viewRepo := NewViewRepository(db)

	testViews := []dbmodels.View{
		{
			UserID:  1,
			VideoID: 1,
			Key:     "1:1",
		},
		{
			UserID:  2,
			VideoID: 1,
			Key:     "2:1",
		},
		{
			UserID:  3,
			VideoID: 1,
			Key:     "3:1",
		},
	}

	for i := range testViews {
		t.Run(fmt.Sprintf("CreateView_Index:%d", i), func(t *testing.T) {
			if err := db.Create(&testViews[i]).Error; err != nil {
				t.Error(err)
			}
		})
	}

	t.Run("GetViewsByVideoID", func(t *testing.T) {
		var (
			views *dto.ViewList
			err   error
		)

		if views, err = viewRepo.GetViewsByVideoID(context.Background(), 1, -1, 0); err != nil {
			t.Error(err)
		}

		if len(views.Views) != len(testViews) {
			t.Error("view lengths are not matched")
		}
	})
}
