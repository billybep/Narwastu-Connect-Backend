package admin

import "app/internal/organization"

type Service interface {
	Create(org *organization.Organization) error
	GetAll() ([]organization.Organization, error)
	GetByID(id uint) (*organization.Organization, error)
	Update(org *organization.Organization) error
	DeleteOrganization(id uint) error
	UpdateProfilePic(id uint, url string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(org *organization.Organization) error {
	return s.repo.Create(org)
}

func (s *service) GetAll() ([]organization.Organization, error) {
	return s.repo.GetAll()
}

func (s *service) GetByID(id uint) (*organization.Organization, error) {
	return s.repo.GetByID(id)
}

func (s *service) Update(org *organization.Organization) error {
	return s.repo.Update(org)
}

func (s *service) DeleteOrganization(id uint) error {
	return s.repo.SoftDelete(id)
}

func (s *service) UpdateProfilePic(id uint, url string) error {
	return s.repo.UpdateProfilePic(id, url)
}
