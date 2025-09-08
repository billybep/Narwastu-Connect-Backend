package verse

import (
	"app/pkg/repository"
)

type VerseRepository struct{}

func NewVerseRepository() *VerseRepository {
	return &VerseRepository{}
}

func (r *VerseRepository) Create(v *Verse) error {
	return repository.DB.Create(v).Error
}

func (r *VerseRepository) GetLatestVerse() (*Verse, error) {
	var verse Verse
	err := repository.DB.
		Order("id DESC").
		Preload("Comments").
		First(&verse).Error
	return &verse, err
}
