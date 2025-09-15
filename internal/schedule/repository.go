package schedule

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]ServiceSchedule, error)
	FindByDate(date time.Time) (*ServiceSchedule, error)
	FindLatest() (*ServiceSchedule, error)
	Create(s *ServiceSchedule) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]ServiceSchedule, error) {
	var schedules []ServiceSchedule
	err := r.db.Order("date desc").Find(&schedules).Error
	return schedules, err
}

func (r *repository) FindByDate(date time.Time) (*ServiceSchedule, error) {
	var schedule ServiceSchedule
	err := r.db.Where("date = ?", date).First(&schedule).Error
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *repository) FindLatest() (*ServiceSchedule, error) {
	var schedule ServiceSchedule
	err := r.db.Order("date desc").First(&schedule).Error
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *repository) Create(s *ServiceSchedule) error {
	return r.db.Create(s).Error
}
