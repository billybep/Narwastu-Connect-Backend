package auth

import "time"

type User struct {
	ID         uint    `gorm:"primaryKey"`
	Provider   string  `gorm:"index;not null"` // "google"
	ProviderID string  `gorm:"index;not null"` // google user id
	Email      *string `gorm:"uniqueIndex"`
	Name       *string
	AvatarURL  *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
