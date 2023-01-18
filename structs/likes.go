package structs

import "time"

type Likes struct {
	Id        int       `json:"id"`
	ComicID   string    `json:"comic_id"`
	UserId    string    `json:"user_id"`
	Like      bool      `json:"like"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
