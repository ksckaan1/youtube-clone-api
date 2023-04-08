package gormadp

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"youtube-clone/pkg/repository/dto"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
)

var _ UserRepositoryInterface = (*UserRepository)(nil)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) CreateUser(ctx context.Context, um *dto.User) error {
	var (
		dbUser dbmodels.User
		err    error
	)

	dbUser.From(um)

	if err = u.db.WithContext(ctx).Create(&dbUser).Error; err != nil {
		return fmt.Errorf("create user : %w", err)
	}

	*um = *dbUser.To()

	return nil
}

func (u UserRepository) GetUserByEmail(ctx context.Context, email string) (*dto.User, error) {
	var (
		dbUser dbmodels.User
		err    error
	)

	if err = u.db.WithContext(ctx).
		Where("email = ?", email).
		First(&dbUser).
		Error; err != nil {
		return nil, fmt.Errorf("get user by email : %w", err)
	}

	return dbUser.To(), nil
}

func (u UserRepository) GetUserByID(ctx context.Context, id uint) (*dto.User, error) {
	var (
		dbUser dbmodels.User
		err    error
	)

	if err = u.db.WithContext(ctx).
		Where("id = ?", id).
		First(&dbUser).
		Error; err != nil {
		return nil, fmt.Errorf("get user by id : %w", err)
	}

	return dbUser.To(), nil
}
