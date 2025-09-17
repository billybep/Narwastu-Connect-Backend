package organization

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Organization, error)
	FindByCategory(category string) ([]Organization, error)
	Create(org Organization) error
	Update(id uint, org Organization) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Organization, error) {
	var orgs []Organization
	err := r.db.Find(&orgs).Error
	return orgs, err
}

func (r *repository) FindByCategory(category string) ([]Organization, error) {
	var orgs []Organization
	err := r.db.Where("category = ?", category).Find(&orgs).Error
	return orgs, err
}

func (r *repository) Create(org Organization) error {
	return r.db.Create(&org).Error
}

func (r *repository) Update(id uint, org Organization) error {
	return r.db.Model(&Organization{}).Where("id = ?", id).Updates(org).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Organization{}, id).Error
}
