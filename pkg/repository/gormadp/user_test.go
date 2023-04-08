package gormadp

import (
	"context"
	"fmt"
	"testing"
	"time"
	"youtube-clone/pkg/repository/dto"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db := NewTestDBConnection(t)
	userRepo := NewUserRepository(db)

	testUsers := []dto.User{
		{
			Email:    "abc@gmail.com",
			Password: "asdf1234",
			FullName: "abc abcoğlu",
		},
		{
			Email:    "def@gmail.com",
			Password: "asdf1234",
			FullName: "abc abcoğlu",
		},
		{
			Email:    "ghi@gmail.com",
			Password: "asdf1234",
			FullName: "abc abcoğlu",
		},
	}

	for i := range testUsers {
		t.Run(fmt.Sprintf("CreateUser_Index:%d", i), func(t *testing.T) {
			if err := userRepo.CreateUser(context.Background(), &testUsers[i]); err != nil {
				t.Error(err)
			}
		})
	}

	t.Run("ErrorTesting", func(t *testing.T) {
		t.Run("DuplicateEmail", func(t *testing.T) {
			dpl := dto.User{
				Email:    "ghi@gmail.com",
				Password: "asdf1234",
				FullName: "abc abcoğlu",
			}

			if err := userRepo.CreateUser(context.Background(), &dpl); err == nil {
				t.Error("expected: UNIQUE constraint failed")
			}
		})

		t.Run("ContextTimeout", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*0)
			defer cancel()

			if err := userRepo.CreateUser(ctx, &dto.User{}); err == nil {
				t.Error("expected: context deadline exceed")
			}
		})

		t.Run("ContextCancel", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			cancel()

			if err := userRepo.CreateUser(ctx, &dto.User{}); err == nil {
				t.Error("expected: context canceled")
			}
		})
	})
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	db := NewTestDBConnection(t)
	userRepo := NewUserRepository(db)

	testUser := dbmodels.User{
		Email:    "abc@gmail.com",
		Password: "asdf1234",
		FullName: "abc abcoğlu",
	}

	t.Run("CreateUser", func(t *testing.T) {
		if err := db.Create(&testUser).Error; err != nil {
			t.Error(err)
		}
	})

	t.Run("GetUserByEmail", func(t *testing.T) {
		if _, err := userRepo.GetUserByEmail(context.Background(), testUser.Email); err != nil {
			t.Error(err)
		}
	})

	t.Run("ErrorTesting", func(t *testing.T) {
		t.Run("NotExistingUser", func(t *testing.T) {
			if _, err := userRepo.GetUserByEmail(context.Background(), "abc"); err == nil {
				t.Error("expected: record not found")
			}
		})

		t.Run("ContextTimeout", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*0)
			defer cancel()

			if _, err := userRepo.GetUserByEmail(ctx, testUser.Email); err == nil {
				t.Error("expected: context deadline exceed")
			}
		})

		t.Run("ContextCancel", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			cancel()

			if _, err := userRepo.GetUserByEmail(ctx, testUser.Email); err == nil {
				t.Error("expected: context canceled")
			}
		})
	})
}

func TestUserRepository_GetUserByID(t *testing.T) {
	db := NewTestDBConnection(t)
	userRepo := NewUserRepository(db)

	testUser := dbmodels.User{
		Email:    "abc@gmail.com",
		Password: "asdf1234",
		FullName: "abc abcoğlu",
	}

	t.Run("CreateUser", func(t *testing.T) {
		if err := db.Create(&testUser).Error; err != nil {
			t.Error(err)
		}
	})

	t.Run("GetUserByID", func(t *testing.T) {
		if _, err := userRepo.GetUserByID(context.Background(), testUser.ID); err != nil {
			t.Error(err)
		}
	})

	t.Run("ErrorTesting", func(t *testing.T) {
		t.Run("NotExistingUser", func(t *testing.T) {
			if _, err := userRepo.GetUserByID(context.Background(), 555); err == nil {
				t.Error("expected: record not found")
			}
		})

		t.Run("ContextTimeout", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*0)
			defer cancel()

			if _, err := userRepo.GetUserByID(ctx, testUser.ID); err == nil {
				t.Error("expected: context deadline exceed")
			}
		})

		t.Run("ContextCancel", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			cancel()

			if _, err := userRepo.GetUserByID(ctx, testUser.ID); err == nil {
				t.Error("expected: context canceled")
			}
		})
	})
}
