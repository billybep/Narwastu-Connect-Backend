package wartajemaat

import (
	"mime/multipart"
	"time"

	supabase "app/internal/storage"

	"gorm.io/gorm"
)

type Service struct {
	db      *gorm.DB
	storage *supabase.Client
}

func NewService(db *gorm.DB, storage *supabase.Client) *Service {
	return &Service{db: db, storage: storage}
}

func (s *Service) GetLatest() (*WartaJemaat, error) {
	var warta WartaJemaat
	err := s.db.Order("created_at DESC").Limit(1).First(&warta).Error
	if err != nil {
		return nil, err
	}
	return &warta, nil
}

func (s *Service) UploadWarta(file multipart.File, fileHeader *multipart.FileHeader) (*WartaJemaat, error) {
	// Upload ke Supabase
	url, err := s.storage.UploadWartaJemaat(file, fileHeader)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	record := &WartaJemaat{
		Title: fileHeader.Filename,
		URL:   url,
		Year:  now.Year(),
		Month: int(now.Month()),
	}

	if err := s.db.Create(record).Error; err != nil {
		return nil, err
	}

	return record, nil
}
