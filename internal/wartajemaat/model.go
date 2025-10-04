package wartajemaat

import "time"

type WartaJemaat struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Year      int       `json:"year"`
	Month     int       `json:"month"`
	CreatedAt time.Time `json:"created_at"`
}
