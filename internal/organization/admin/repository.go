package admin

import (
	"app/internal/organization"

	"gorm.io/gorm"
)

type Repository interface {
	Create(org *organization.Organization) error
	GetAll() ([]organization.Organization, error)
	GetByID(id uint) (*organization.Organization, error)
	Update(org *organization.Organization) error
	SoftDelete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(org *organization.Organization) error {
	return r.db.Create(org).Error
}

func (r *repository) GetAll() ([]organization.Organization, error) {
	var orgs []organization.Organization
	err := r.db.Find(&orgs).Error
	return orgs, err
}

func (r *repository) GetByID(id uint) (*organization.Organization, error) {
	var org organization.Organization
	err := r.db.First(&org, id).Error
	return &org, err
}

func (r *repository) Update(org *organization.Organization) error {
	return r.db.Save(org).Error
}

func (r *repository) SoftDelete(id uint) error {
	return r.db.Delete(&organization.Organization{}, id).Error
}
