package structs

import "time"

type Comment struct {
	Id          int       `json:"id"`
	ComicId     int       `json:"comic_id"`
	UserId      int       `json:"user_id"`
	ReferenceId int       `json:"reference_id"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
