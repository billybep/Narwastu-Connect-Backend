package sermon

import (
	"time"

	"gorm.io/gorm"
)

type Sermon struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Title        string         `gorm:"size:200;not null" json:"title"`
	MainVerse    string         `gorm:"size:100;not null" json:"main_verse"`
	Preacher     string         `gorm:"size:100;not null" json:"preacher"`
	Date         time.Time      `gorm:"not null" json:"date"`
	Content      string         `gorm:"type:text;not null" json:"content"`
	ImageURL     string         `json:"image_url"`
	ProfileImage string         `json:"profile_image"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
