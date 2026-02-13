package models

import "time"

// Blog represents a blog post
type Blog struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ViewCount   int       `json:"view_count"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
}
