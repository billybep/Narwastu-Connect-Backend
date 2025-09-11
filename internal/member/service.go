package member

import (
	"time"
)

type Service interface {
	GetWeeklyBirthdays() ([]Member, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetWeeklyBirthdays() ([]Member, error) {
	now := time.Now()
	// range minggu: Minggu (Sunday) - Sabtu (Saturday)
	weekday := int(now.Weekday()) // 0=Sunday
	start := now.AddDate(0, 0, -weekday)
	end := start.AddDate(0, 0, 6)

	return s.repo.FindBirthdaysInRange(start, end)
}
