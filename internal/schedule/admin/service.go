package admin

import "app/internal/schedule"

type Service interface {
	CreateSchedule(s *schedule.ServiceSchedule) error
	GetSchedules() ([]schedule.ServiceSchedule, error)
	GetScheduleByID(id uint) (*schedule.ServiceSchedule, error)
	UpdateSchedule(id uint, data *schedule.ServiceSchedule) error
	DeleteSchedule(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateSchedule(sc *schedule.ServiceSchedule) error {
	return s.repo.Create(sc)
}

func (s *service) GetSchedules() ([]schedule.ServiceSchedule, error) {
	return s.repo.GetAll()
}

func (s *service) GetScheduleByID(id uint) (*schedule.ServiceSchedule, error) {
	return s.repo.GetByID(id)
}

func (s *service) UpdateSchedule(id uint, data *schedule.ServiceSchedule) error {
	return s.repo.Update(id, data)
}

func (s *service) DeleteSchedule(id uint) error {
	return s.repo.SoftDelete(id)
}
