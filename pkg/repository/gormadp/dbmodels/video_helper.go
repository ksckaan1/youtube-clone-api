package dbmodels

import (
	"gorm.io/gorm"
	"youtube-clone/pkg/repository/dto"
)

func (v *Video) From(m *dto.Video) {
	*v = Video{
		Model: gorm.Model{
			ID:        m.ID,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
		Title:       m.Title,
		Description: m.Description,
		URL:         m.URL,
		Thumbnail:   m.Thumbnail,
		UserID:      m.UserID,
	}
}

func (v *Video) To() *dto.Video {
	return &dto.Video{
		StdModel: dto.StdModel{
			ID:        v.UserID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		},
		Title:       v.Title,
		Description: v.Description,
		URL:         v.URL,
		Thumbnail:   v.Thumbnail,
		UserID:      v.UserID,
	}
}
