package member

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindBirthdaysInRange(start, end time.Time) ([]Member, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindBirthdaysInRange(start, end time.Time) ([]Member, error) {
	var members []Member
	// cek hanya bulan & hari, abaikan tahun
	err := r.db.
		Preload("Site"). // 👈 preload relasi site
		Where(`
		EXTRACT(MONTH FROM date_of_birth) = ? 
		AND EXTRACT(DAY FROM date_of_birth) BETWEEN ? AND ?
	`, int(start.Month()), start.Day(), end.Day()).Find(&members).Error
	return members, err
}
