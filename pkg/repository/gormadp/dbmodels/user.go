package dbmodels

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Password string // `gorm:"type:varchar(32)"` 32 karakterlik metin ayır
	FullName string
}
