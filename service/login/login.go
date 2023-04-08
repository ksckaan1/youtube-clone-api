package login

import (
	"context"
	"fmt"
	"youtube-clone/pkg/jwt"
	"youtube-clone/pkg/jwt/golang_jwt"
	"youtube-clone/pkg/passwd"
	"youtube-clone/pkg/passwd/bcrypt"
	"youtube-clone/pkg/repository/dto"
	"youtube-clone/pkg/repository/gormadp"
	"youtube-clone/pkg/validator"
)

var _ Interface = (*Login)(nil)

type Login struct {
	repo *gormadp.Repository
	pass passwd.Interface
	jwt  jwt.Interface
}

func New(repo *gormadp.Repository, secret string, expHour int64) *Login {
	bc := bcrypt.New(secret, 10)

	j := golang_jwt.New(secret, expHour)

	return &Login{repo: repo, pass: bc, jwt: j}
}

func (l *Login) Login(ctx context.Context, email, password string) (string, error) {
	var (
		user *dto.User
		err  error
	)

	emailVld := validator.IsEmailValid{Email: email}
	passwordVld := validator.IsInRange{
		Text: password,
		Min:  6,
		Max:  32,
	}

	if err = validator.ValidateAll(&emailVld, &passwordVld); err != nil {
		return "", fmt.Errorf("login : %w", err)
	}

	if user, err = l.repo.UserRepo.GetUserByEmail(ctx, email); err != nil {
		return "", fmt.Errorf("login : %w", err)
	}

	if err = l.pass.Compare(user.Password, password); err != nil {
		return "", fmt.Errorf("login : %w", ErrPasswordMatch)
	}

	var token string

	if token, err = l.jwt.Generate(user.ID); err != nil {
		return "", fmt.Errorf("login : %w", err)
	}

	return token, nil
}
