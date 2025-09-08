package verse

import (
	"errors"

	"app/pkg/repository"

	"gorm.io/gorm"
)

type VerseService struct {
	repo *VerseRepository
}

func NewVerseService(repo *VerseRepository) *VerseService {
	return &VerseService{repo: repo}
}

// Create new verse
func (s *VerseService) CreateVerse(v *Verse) error {
	return s.repo.Create(v)
}

// Get latest verse
func (s *VerseService) GetLatestVerse() (*Verse, error) {
	return s.repo.GetLatestVerse()
}

// Like a verse (only once per user)
func (s *VerseService) LikeVerse(userID, verseID uint) error {
	// cek verse ada
	var v Verse
	if err := repository.DB.First(&v, verseID).Error; err != nil {
		return errors.New("verse not found")
	}

	// cek apakah user sudah like
	var existing VerseLike
	if err := repository.DB.Where("verse_id = ? AND user_id = ?", verseID, userID).First(&existing).Error; err == nil {
		return errors.New("already liked")
	}

	// insert like baru
	like := VerseLike{VerseID: verseID, UserID: userID}
	if err := repository.DB.Create(&like).Error; err != nil {
		return err
	}

	// increment counter
	return repository.DB.Model(&v).Update("likes", gorm.Expr("likes + ?", 1)).Error
}

// Share a verse
func (s *VerseService) ShareVerse(id uint) error {
	var v Verse
	if err := repository.DB.First(&v, id).Error; err != nil {
		return errors.New("verse not found")
	}
	return repository.DB.Model(&v).Update("shares", gorm.Expr("shares + ?", 1)).Error
}

// Add comment
func (s *VerseService) AddComment(id uint, content string, user string) error {
	var v Verse
	if err := repository.DB.First(&v, id).Error; err != nil {
		return errors.New("verse not found")
	}
	comment := Comment{
		VerseID: id,
		User:    user,
		Content: content,
	}
	return repository.DB.Create(&comment).Error
}
