package schedule

import "time"

type Service interface {
	GetAll() ([]ServiceSchedule, error)
	GetByDate(date time.Time) (*ServiceSchedule, error)
	GetLatest() (*ServiceSchedule, error)
	Create(schedule *ServiceSchedule) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]ServiceSchedule, error) {
	return s.repo.FindAll()
}

func (s *service) GetByDate(date time.Time) (*ServiceSchedule, error) {
	return s.repo.FindByDate(date)
}

func (s *service) GetLatest() (*ServiceSchedule, error) {
	return s.repo.FindLatest()
}

func (s *service) Create(schedule *ServiceSchedule) error {
	return s.repo.Create(schedule)
}
