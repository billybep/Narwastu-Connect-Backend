package verse

import "time"

type Verse struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	VerseReference string    `json:"verseReference"`
	VerseText      string    `json:"verseText"`
	Likes          int       `json:"likes"`
	Shares         int       `json:"shares"`
	CreatedAt      time.Time `json:"createdAt"`
	Comments       []Comment `json:"comments"` // relasi one-to-many
}

type VerseLike struct {
	ID        uint `gorm:"primaryKey"`
	VerseID   uint `gorm:"index;not null"`
	UserID    uint `gorm:"index;not null"`
	CreatedAt time.Time

	// Optional: relasi GORM
	Verse Verse `gorm:"foreignKey:VerseID"`
}

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	VerseID   uint      `json:"verseId"`
	User      string    `json:"user"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
