package dbmodels

import (
	"youtube-clone/pkg/enum/liketype"
	"youtube-clone/pkg/repository/dto"
)

func (l *Like) From(m *dto.Like) {
	*l = Like{
		UserID:  m.UserID,
		VideoID: m.VideoID,
		Type:    uint(m.Type),
		Key:     m.Key,
	}
}

func (l *Like) To() *dto.Like {
	return &dto.Like{
		UserID:  l.UserID,
		VideoID: l.VideoID,
		Type:    liketype.LikeType(l.Type),
		Key:     l.Key,
	}
}
