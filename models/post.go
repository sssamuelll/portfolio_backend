package models

type Post struct {
	ID          uint     `gorm:"primaryKey"`
	Image       string   `json:"image"`
	Name        string   `json:"name" gorm:"not null"`
	Description string   `json:"description"`
	Category    string   `json:"category" gorm:"not null"`
	Tags        []string `gorm:"-" json:"tags"`
	TagsJSON    string   `json:"-" gorm:"column:tags"`
	Media       []string `gorm:"-" json:"media"`
	MediaJSON   string   `json:"-" gorm:"column:media"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Link        string   `json:"link"`
}
