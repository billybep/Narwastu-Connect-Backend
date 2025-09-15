package member

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindBirthdaysInRange(start, end time.Time) ([]Member, error)
	CreateMember(member *Member) error
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
	`, int(start.Month()), start.Day(), end.Day()).
		Order("EXTRACT(DAY FROM date_of_birth) ASC").Find(&members).Error
	return members, err
}

// Create Member
func (r *repository) CreateMember(member *Member) error {
	return r.db.Create(member).Error
}
