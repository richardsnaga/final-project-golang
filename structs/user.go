package structs

import "time"

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
