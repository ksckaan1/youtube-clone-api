package dbmodels

import "youtube-clone/pkg/repository/dto"

func (s *Subscribe) From(m *dto.Subscribe) {
	*s = Subscribe{
		UserID:    m.UserID,
		ChannelID: m.ChannelID,
	}
}

func (s *Subscribe) To() *dto.Subscribe {
	return &dto.Subscribe{
		UserID:    s.UserID,
		ChannelID: s.ChannelID,
	}
}
