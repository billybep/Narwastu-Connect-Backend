package admin

import (
	"app/internal/event"
	"errors"
	"time"
)

type Service interface {
	CreateEvent(title, location, description, imageUrl string, dateTime time.Time) error
	GetAllEvents() ([]event.Event, error)
	GetEventByID(id uint) (*event.Event, error)
	UpdateEvent(id uint, title, location, description, imageUrl string, dateTime time.Time) error
	DeleteEvent(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateEvent(title, location, description, imageUrl string, dateTime time.Time) error {
	if title == "" {
		return errors.New("title is required")
	}

	event := &event.Event{
		Title:       title,
		DateTime:    dateTime,
		Location:    location,
		Description: description,
		ImageURL:    imageUrl,
	}

	return s.repo.CreateEvent(event)
}

func (s *service) GetAllEvents() ([]event.Event, error) {
	return s.repo.GetAllEvents()
}

func (s *service) GetEventByID(id uint) (*event.Event, error) {
	return s.repo.GetEventByID(id)
}

func (s *service) UpdateEvent(id uint, title, location, description, imageUrl string, dateTime time.Time) error {
	eventData, err := s.repo.GetEventByID(id)
	if err != nil {
		return errors.New("event not found")
	}

	if title != "" {
		eventData.Title = title
	}
	if !dateTime.IsZero() {
		eventData.DateTime = dateTime
	}
	if location != "" {
		eventData.Location = location
	}
	if description != "" {
		eventData.Description = description
	}
	if imageUrl != "" {
		eventData.ImageURL = imageUrl
	}

	return s.repo.UpdateEvent(eventData)
}

func (s *service) DeleteEvent(id uint) error {
	return s.repo.SoftDeleteEvent(id)
}
