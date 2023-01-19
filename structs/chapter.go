package structs

import "time"

type Chapter struct {
	Id            int       `json:"id"`
	ComicID       int    `json:"comic_id"`
	ChapterNumber int       `json:"chapter_number"`
	ImageUrl      string    `json:"image_url"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}
