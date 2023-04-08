package login

import (
	"context"
	"testing"
	"youtube-clone/pkg/repository/gormadp"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
)

var (
	secret  = "verysecretkey"
	expHour = int64(24)
)

func TestLogin_Login(t *testing.T) {
	db := gormadp.NewTestDBConnection(t)

	repo := gormadp.NewRepository(db)

	testUser := dbmodels.User{
		Email:    "me@kaanksc.com",
		Password: "$2a$10$XQSil8wLLOjbhsuxnhsBdOH6Z8dXszXA9b2ELKLZeuN.DG13aTHSi", // asdf1234
		FullName: "Kaan Kuscu",
	}

	t.Run("CreateUser", func(t *testing.T) {
		if err := db.Create(&testUser).Error; err != nil {
			t.Error(err)
		}
	})

	loginService := New(repo, secret, expHour)

	t.Run("Login", func(t *testing.T) {
		if _, err := loginService.Login(context.Background(), testUser.Email, "asdf1234"); err != nil {
			t.Error(err)
		}
	})

	t.Run("ErrorTesting", func(t *testing.T) {
		t.Run("WrongEmail", func(t *testing.T) {
			if _, err := loginService.Login(context.Background(), "abc@gmail.com", "asdf1234"); err == nil {
				t.Error("expected: record not found")
			}
		})

		t.Run("WrongPassword", func(t *testing.T) {
			if _, err := loginService.Login(context.Background(), testUser.Email, "abc123"); err == nil {
				t.Error("expected: wrong password")
			}
		})
	})
}
