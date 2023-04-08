package dto

import "youtube-clone/pkg/enum/liketype"

type Like struct {
	UserID  uint              `json:"user_id"`
	VideoID uint              `json:"video"`
	Type    liketype.LikeType `json:"type"`
	Key     string            `json:"key"`
}

type LikeList struct {
	Limit  int     `json:"limit"`
	Offset int     `json:"offset"`
	Count  int64   `json:"count"`
	Likes  []*Like `json:"likes"`
}
