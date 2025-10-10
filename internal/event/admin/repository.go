package admin

import (
	"app/internal/event"

	"gorm.io/gorm"
)

type Repository interface {
	CreateEvent(e *event.Event) error
	GetAllEvents() ([]event.Event, error)
	GetEventByID(id uint) (*event.Event, error)
	UpdateEvent(e *event.Event) error
	SoftDeleteEvent(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateEvent(e *event.Event) error {
	return r.db.Create(e).Error
}

func (r *repository) GetAllEvents() ([]event.Event, error) {
	var events []event.Event
	err := r.db.Order("date_time DESC").Find(&events).Error
	return events, err
}

func (r *repository) GetEventByID(id uint) (*event.Event, error) {
	var e event.Event
	err := r.db.First(&e, id).Error
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *repository) UpdateEvent(e *event.Event) error {
	return r.db.Save(e).Error
}

func (r *repository) SoftDeleteEvent(id uint) error {
	return r.db.Delete(&event.Event{}, id).Error
}
