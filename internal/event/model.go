package event

import "time"

type Event struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	DateTime    time.Time `json:"dateTime"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	ImageURL    string    `json:"imageUrl"`
}
