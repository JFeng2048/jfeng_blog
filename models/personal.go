package models

// PersonalInfo represents personal information
type PersonalInfo struct {
	Name        string            `json:"name"`
	Title       string            `json:"title"`
	Bio         string            `json:"bio"`
	Email       string            `json:"email"`
	Location    string            `json:"location"`
	Avatar      string            `json:"avatar"`
	SocialLinks map[string]string `json:"social_links"`
	Skills      []string          `json:"skills"`
	Education   []Education       `json:"education"`
	Experience  []Experience      `json:"experience"`
}

// Education represents educational background
type Education struct {
	School     string `json:"school"`
	Degree     string `json:"degree"`
	Field      string `json:"field"`
	StartYear  int    `json:"start_year"`
	EndYear    int    `json:"end_year"`
}

// Experience represents work experience
type Experience struct {
	Company     string   `json:"company"`
	Position    string   `json:"position"`
	Description string   `json:"description"`
	StartYear   int      `json:"start_year"`
	EndYear     *int     `json:"end_year,omitempty"`
	Technologies []string `json:"technologies"`
}
