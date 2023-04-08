package dto

import "time"

type View struct {
	CreatedAt time.Time `json:"created_at"`
	UserID    uint      `json:"user_id"`
	VideoID   uint      `json:"video_id"`
	Key       string    `json:"key"`
}

type ViewList struct {
	Limit  int     `json:"limit"`
	Offset int     `json:"offset"`
	Count  int64   `json:"count"`
	Views  []*View `json:"views"`
}
