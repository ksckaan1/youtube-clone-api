package dbmodels

import "time"

type View struct {
	CreatedAt time.Time
	UserID    uint
	VideoID   uint
	Key       string `gorm:"uniqueIndex"`
}
