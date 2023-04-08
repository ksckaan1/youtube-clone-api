package dto

type Video struct {
	StdModel
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Thumbnail   string `json:"thumbnail"`
	UserID      uint   `json:"user_id"`
}

type VideoList struct {
	Limit  int      `json:"limit"`
	Offset int      `json:"offset"`
	Count  int64    `json:"count"`
	Videos []*Video `json:"videos"`
}
