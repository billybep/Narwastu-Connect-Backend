package admin

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(email, password string) (*Admin, error)
	GetAllAdmins() ([]Admin, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Login(email, password string) (*Admin, error) {
	admin, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}
	return admin, nil
}

func (s *service) GetAllAdmins() ([]Admin, error) {
	return s.repo.GetAll()
}
