package models

import "time"

// Project represents a project in the portfolio
type Project struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Technologies []string   `json:"technologies"`
	GithubURL    string     `json:"github_url"`
	LiveURL      string     `json:"live_url"`
	ImageURL     string     `json:"image_url"`
	StartDate    time.Time  `json:"start_date"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	Status       string     `json:"status"` // "active", "completed", "archived"
}
