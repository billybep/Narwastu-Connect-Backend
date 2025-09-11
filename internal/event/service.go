package event

import (
	"log"
	"time"
)

type Service interface {
	GetEvents() ([]Event, error)
	GetWeeklyEvents(now time.Time) ([]Event, error)
	GetEventByID(id uint) (*Event, error)
	CreateEvent(event *Event) error
	UpdateEvent(event *Event) error
	DeleteEvent(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetEvents() ([]Event, error) {
	return s.repo.FindAll()
}

func (s *service) GetWeeklyEvents(now time.Time) ([]Event, error) {
	now = now.UTC()

	weekday := int(now.Weekday()) // Minggu = 0, Senin = 1, dst.

	// Mulai minggu: Minggu (00:00 UTC)
	start := time.Date(now.Year(), now.Month(), now.Day()-weekday, 0, 0, 0, 0, time.UTC)

	// Akhir minggu: Minggu depan (23:59:59 UTC)
	end := start.AddDate(0, 0, 8).Add(-time.Nanosecond)

	log.Printf("Weekly range: %v - %v", start, end)

	return s.repo.FindWeekly(start, end)
}

func (s *service) GetEventByID(id uint) (*Event, error) {
	return s.repo.FindByID(id)
}

func (s *service) CreateEvent(event *Event) error {
	return s.repo.Create(event)
}

func (s *service) UpdateEvent(event *Event) error {
	return s.repo.Update(event)
}

func (s *service) DeleteEvent(id uint) error {
	return s.repo.Delete(id)
}
