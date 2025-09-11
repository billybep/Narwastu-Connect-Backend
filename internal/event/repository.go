package event

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindWeekly(start, end time.Time) ([]Event, error)
	FindByID(id uint) (*Event, error)
	Create(event *Event) error
	Update(event *Event) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Event, error) {
	var events []Event
	err := r.db.Find(&events).Error
	return events, err
}

func (r *repository) FindWeekly(start, end time.Time) ([]Event, error) {
	var events []Event
	err := r.db.Where("date_time BETWEEN ? AND ?", start, end).Order("date_time ASC").Find(&events).Error
	return events, err
}

func (r *repository) FindByID(id uint) (*Event, error) {
	var event Event
	err := r.db.First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *repository) Create(event *Event) error {
	return r.db.Create(event).Error
}

func (r *repository) Update(event *Event) error {
	return r.db.Save(event).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Event{}, id).Error
}
