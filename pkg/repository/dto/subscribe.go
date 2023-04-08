package dto

type Subscribe struct {
	UserID    uint
	ChannelID uint
	Key       string
}

type SubscribeList struct {
	Limit      int
	Offset     int
	Count      int64
	Subscribes []*Subscribe
}
