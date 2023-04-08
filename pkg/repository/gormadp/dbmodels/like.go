package dbmodels

type Like struct {
	UserID  uint
	VideoID uint
	Type    uint
	Key     string `gorm:"uniqueIndex"`
}
