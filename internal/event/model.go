package event

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	DateTime    time.Time      `json:"dateTime"`
	Location    string         `json:"location"`
	Description string         `json:"description"`
	ImageURL    string         `json:"imageUrl"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
