package member

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindBirthdaysInRange(start, end time.Time) ([]Member, error)
	CreateMember(member *Member) error
	FindAllMembers() ([]Member, error)
	FindMemberByID(id uint) (*Member, error)
	UpdateMember(member *Member) error
	DeleteMember(id uint) error
	UpdateMemberAvatar(memberID uint, avatarURL string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAllMembers() ([]Member, error) {
	var members []Member
	err := r.db.
		Preload("Site").
		Order("full_name ASC").
		Find(&members).Error
	return members, err
}

func (r *repository) FindMemberByID(id uint) (*Member, error) {
	var member Member
	err := r.db.
		Preload("Site").
		First(&member, id).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *repository) UpdateMember(member *Member) error {
	return r.db.Save(member).Error
}

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

// Soft Delete Member
func (r *repository) DeleteMember(id uint) error {
	// Soft delete otomatis karena model Member punya gorm.DeletedAt
	return r.db.Delete(&Member{}, id).Error
}

func (r *repository) UpdateMemberAvatar(memberID uint, avatarURL string) error {
	return r.db.Model(&Member{}).
		Where("id = ?", memberID).
		Update("photo_url", avatarURL).Error
}
