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

// func (r *repository) FindBirthdaysInRange(start, end time.Time) ([]Member, error) {
// 	var members []Member
// 	// cek hanya bulan & hari, abaikan tahun
// 	err := r.db.
// 		Preload("Site"). // 👈 preload relasi site
// 		Where(`
// 		EXTRACT(MONTH FROM date_of_birth) = ?
// 		AND EXTRACT(DAY FROM date_of_birth) BETWEEN ? AND ?
// 	`, int(start.Month()), start.Day(), end.Day()).
// 		Order("EXTRACT(DAY FROM date_of_birth) ASC").Find(&members).Error
// 	return members, err
// }

func (r *repository) FindBirthdaysInRange(start, end time.Time) ([]Member, error) {
	var members []Member

	if start.Month() == end.Month() {
		// minggu masih di bulan yang sama
		err := r.db.
			Preload("Site").
			Where(`
				EXTRACT(MONTH FROM date_of_birth) = ? 
				AND EXTRACT(DAY FROM date_of_birth) BETWEEN ? AND ?
			`, int(start.Month()), start.Day(), end.Day()).
			Order("EXTRACT(DAY FROM date_of_birth) ASC").
			Find(&members).Error
		return members, err
	} else {
		// minggu melewati batas bulan (misal: 28 Sep – 4 Okt, atau 29 Des – 4 Jan)
		err := r.db.
			Preload("Site").
			Where(`
				(EXTRACT(MONTH FROM date_of_birth) = ? AND EXTRACT(DAY FROM date_of_birth) >= ?)
				OR
				(EXTRACT(MONTH FROM date_of_birth) = ? AND EXTRACT(DAY FROM date_of_birth) <= ?)
			`, int(start.Month()), start.Day(), int(end.Month()), end.Day()).
			Order("EXTRACT(MONTH FROM date_of_birth), EXTRACT(DAY FROM date_of_birth) ASC").
			Find(&members).Error
		return members, err
	}
}

// Create Member
func (r *repository) CreateMember(member *Member) error {
	return r.db.Create(member).Error
}
