package models

import "time"

type BlogPost struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Excerpt   string    `json:"excerpt"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
