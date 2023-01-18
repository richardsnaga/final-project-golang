package structs

import "time"

type Rating struct {
	Id        int       `json:"id"`
	ComicId   int       `json:"comic_id"`
	UserId    int       `json:"user_id"`
	Rate      int       `json:"rate"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type AvgRating struct {
	ComicId int     `json:"comic_id"`
	AvgRate float64 `json:"avg_rate"`
}
