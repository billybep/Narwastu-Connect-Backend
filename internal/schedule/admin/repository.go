package admin

import (
	"app/internal/schedule"

	"gorm.io/gorm"
)

type Repository interface {
	Create(schedule *schedule.ServiceSchedule) error
	GetAll() ([]schedule.ServiceSchedule, error)
	GetByID(id uint) (*schedule.ServiceSchedule, error)
	Update(id uint, data *schedule.ServiceSchedule) error
	SoftDelete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(s *schedule.ServiceSchedule) error {
	return r.db.Create(s).Error
}

func (r *repository) GetAll() ([]schedule.ServiceSchedule, error) {
	var schedules []schedule.ServiceSchedule
	err := r.db.Order("date DESC").Find(&schedules).Error
	return schedules, err
}

func (r *repository) GetByID(id uint) (*schedule.ServiceSchedule, error) {
	var s schedule.ServiceSchedule
	err := r.db.First(&s, id).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *repository) Update(id uint, data *schedule.ServiceSchedule) error {
	return r.db.Model(&schedule.ServiceSchedule{}).Where("id = ?", id).Updates(data).Error
}

func (r *repository) SoftDelete(id uint) error {
	return r.db.Delete(&schedule.ServiceSchedule{}, id).Error
}
