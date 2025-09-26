package verse

import (
	"time"
)

type Verse struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	VerseReference string    `json:"verseReference"`
	VerseText      string    `json:"verseText"`
	Likes          int64     `gorm:"-" json:"likes"` //
	Shares         int       `json:"shares"`
	CreatedAt      time.Time `json:"createdAt"`
}

type VerseLike struct {
	ID        uint `gorm:"primaryKey"`
	VerseID   uint `gorm:"index;not null"`
	UserID    uint `gorm:"index;not null"`
	CreatedAt time.Time

	Verse Verse `gorm:"foreignKey:VerseID;constraint:OnDelete:CASCADE;"`
}
