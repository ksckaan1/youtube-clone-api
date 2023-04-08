package gormadp

import (
	"fmt"
	"gorm.io/gorm"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
)

type Repository struct {
	UserRepo      *UserRepository
	SubscribeRepo *SubscribeRepository
	VideoRepo     *VideoRepository
	ViewRepo      *ViewRepository
	LikeRepo      *LikeRepository
	db            *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:            db,
		UserRepo:      NewUserRepository(db),
		SubscribeRepo: NewSubscribeRepository(db),
		VideoRepo:     NewVideoRepository(db),
		ViewRepo:      NewViewRepository(db),
		LikeRepo:      NewLikeRepository(db),
	}
}

func (r *Repository) AutoMigrate() error {
	if err := r.db.AutoMigrate(
		&dbmodels.User{},
		&dbmodels.Subscribe{},
		&dbmodels.Video{},
		&dbmodels.View{},
	); err != nil {
		return fmt.Errorf("auto migrate : %w", err)
	}

	return nil
}
