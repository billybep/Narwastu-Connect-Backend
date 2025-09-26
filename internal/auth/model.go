package auth

import (
	"app/internal/verse"
	"time"
)

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Provider   string `gorm:"index;not null"` // "google"
	ProviderID string `gorm:"index;not null"` // google user id
	Email      string `gorm:"uniqueIndex" json:"email"`
	Name       *string
	AvatarURL  *string
	CreatedAt  time.Time
	UpdatedAt  time.Time

	// relasi ke verse_likes
	VerseLikes []verse.VerseLike `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"-"`
}
