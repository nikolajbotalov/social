package domain

import "time"

type Post struct {
	ID        string    `json:"id"`
	ChannelID string    `json:"channel_id"`
	AuthorID  string    `json:"author_id"`
	Content   string    `json:"content"`
	ImageURL  string    `json:"image_url,omitempty"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	CreatedAt time.Time `json:"created_at"`
}
