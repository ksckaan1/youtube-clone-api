package video

import "time"

type CreateVideoModel struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}

type GetVideoModel struct {
	Id          uint            `json:"id"`
	CreatedAt   time.Time       `json:"created_at"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	ViewCount   int64           `json:"view_count"`
	LikeCount   int64           `json:"like_count"`
	IsLiked     bool            `json:"is_liked"`
	IsDisliked  bool            `json:"is_disliked"`
	Comments    []*CommentModel `json:"comments"`
}

type CommentModel struct {
	ID        uint              `json:"id"`
	CreatedAt time.Time         `json:"created_at"`
	Message   string            `json:"comment"`
	User      *CommentUserModel `json:"user"`
}

type CommentUserModel struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
}
