package admin

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetByID(id uint) (*Admin, error)
	FindByEmail(email string) (*Admin, error)
	Create(admin *Admin) error
	GetAll() ([]Admin, error)
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetByID(id uint) (*Admin, error) {
	var admin Admin
	if err := r.db.First(&admin, id).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *repository) FindByEmail(email string) (*Admin, error) {
	var admin Admin
	if err := r.db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *repository) Create(admin *Admin) error {
	return r.db.Create(admin).Error
}

func (r *repository) GetAll() ([]Admin, error) {
	var admins []Admin
	if err := r.db.Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Admin{}, id).Error
}
