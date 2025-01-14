package models

type Post struct {
	ID          int    `json:"id" db:"id"`
	Image       string `json:"image" db:"image"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Category    string `json:"category" db:"category"`
	Tags        string `json:"tags" db:"tags"`   // Almacenar como JSON o CSV
	Media       string `json:"media" db:"media"` // Almacenar como JSON o CSV
	StartDate   string `json:"startDate,omitempty" db:"start_date"`
	EndDate     string `json:"endDate,omitempty" db:"end_date"`
	Link        string `json:"link,omitempty" db:"link"`
}
