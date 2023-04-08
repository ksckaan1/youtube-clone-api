package dbmodels

import (
	"youtube-clone/pkg/repository/dto"
)

func (v *View) From(m *dto.View) {
	*v = View{
		CreatedAt: m.CreatedAt,
		UserID:    m.UserID,
		VideoID:   m.VideoID,
		Key:       m.Key,
	}
}

func (v *View) To() *dto.View {
	return &dto.View{
		CreatedAt: v.CreatedAt,
		UserID:    v.UserID,
		VideoID:   v.VideoID,
		Key:       v.Key,
	}
}
