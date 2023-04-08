package dbmodels

import (
	"gorm.io/gorm"
	"youtube-clone/pkg/repository/dto"
)

func (u *User) From(m *dto.User) {
	*u = User{
		Model: gorm.Model{
			ID:        m.ID,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
		Email:    m.Email,
		Password: m.Password,
		FullName: m.FullName,
	}
}

func (u *User) To() *dto.User {
	return &dto.User{
		StdModel: dto.StdModel{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		Email:    u.Email,
		Password: u.Password,
		FullName: u.FullName,
	}
}
