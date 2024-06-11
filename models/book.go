package models

import "time"

type Book struct {
	ID        uint       `json:"id" example:"1"`
	CreatedAt time.Time  `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time  `json:"updated_at" example:"2023-01-02T00:00:00Z"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" example:"2023-01-03T00:00:00Z"`
	Title     string     `json:"title" example:"The Great Gatsby"`
	Author    string     `json:"author" example:"F. Scott Fitzgerald"`
	Year      int        `json:"year" example:"1925"`
}
