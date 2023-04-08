package dbmodels

type Subscribe struct {
	UserID    uint   `gorm:"index"`
	ChannelID uint   `gorm:"index"`
	Key       string `gorm:"uniqueIndex"`
}
