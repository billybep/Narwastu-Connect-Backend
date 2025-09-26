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

	// Ambil 1 verse terbaru
	err := repository.DB.
		Order("id DESC").
		First(&verse).Error
	if err != nil {
		return nil, err
	}

	// Hitung likes dari verse_likes
	var count int64
	err = repository.DB.Model(&VerseLike{}).
		Where("verse_id = ?", verse.ID).
		Count(&count).Error
	if err != nil {
		return nil, err
	}
	verse.Likes = count

	return &verse, nil
}
