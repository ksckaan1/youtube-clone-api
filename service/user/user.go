package user

import (
	"context"
	"fmt"
	"youtube-clone/pkg/passwd"
	"youtube-clone/pkg/passwd/bcrypt"
	"youtube-clone/pkg/repository/dto"
	"youtube-clone/pkg/repository/gormadp"
	"youtube-clone/pkg/validator"
)

var _ Interface = (*User)(nil)

type User struct {
	repo *gormadp.Repository
	pw   passwd.Interface
}

func New(repo *gormadp.Repository, secret string) (*User, error) {
	if repo == nil {
		return nil, fmt.Errorf("new user service : %w", ErrRepositoryShouldBeSetUp)
	}

	pw := bcrypt.New(secret, 10)

	return &User{repo: repo, pw: pw}, nil
}

func (u *User) CreateUser(ctx context.Context, m *CreateUserModel) error {
	var err error

	emailVld := validator.IsEmailValid{Email: m.Email}
	passwdVld := validator.IsInRange{Min: 6, Max: 32, Text: m.Password}

	if err = validator.ValidateAll(&emailVld, &passwdVld); err != nil {
		return fmt.Errorf("create user : %w", err)
	}

	var hashedPW string

	if hashedPW, err = u.pw.Generate(m.Password); err != nil {
		return fmt.Errorf("create user / generate password : %w", err)
	}

	userDTO := dto.User{
		Email:    m.Email,
		Password: hashedPW,
		FullName: m.FullName,
	}

	if err = u.repo.UserRepo.CreateUser(ctx, &userDTO); err != nil {
		return fmt.Errorf("create user / create : %w", err)
	}

	return nil
}

func (u *User) GetUserByID(ctx context.Context, userID uint) (*GetUserModel, error) {
	var (
		user *dto.User
		err  error
	)

	if user, err = u.repo.UserRepo.GetUserByID(ctx, userID); err != nil {
		return nil, fmt.Errorf("get user by id : %w", err)
	}

	return &GetUserModel{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
	}, nil
}
