package structs

import "time"

type Comment struct {
	Id          int       `json:"id"`
	ChapterId   int       `json:"chapter_id"`
	UserId      int       `json:"user_id"`
	ReferenceId int       `json:"reference_id"`
	Comment     string    `json:"comment"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
