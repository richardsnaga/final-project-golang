package structs

import "time"

type Chapters struct {
	Id            int       `json:"id"`
	ComicID       string    `json:"comic_id"`
	ChapterNumber int       `json:"chapter_number"`
	ImageUrl      string    `json:"image_url"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}
