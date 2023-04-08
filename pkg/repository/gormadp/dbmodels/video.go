package dbmodels

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title       string
	Description string
	URL         string
	Thumbnail   string
	UserID      uint `gorm:"index"`
}
