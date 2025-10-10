package admin

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type CreateAdminInput struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type Service interface {
	Login(email, password string) (*Admin, error)
	CreateAdmin(input CreateAdminInput, creator *Admin) (*Admin, error)
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

func (s *service) CreateAdmin(input CreateAdminInput, creator *Admin) (*Admin, error) {
	// hanya system admin yang boleh buat admin baru
	if creator.Role != RoleSystemAdmin {
		return nil, errors.New("unauthorized: only system administrator can create admin")
	}

	// cek email duplikat
	existing, _ := s.repo.FindByEmail(input.Email)
	if existing != nil && existing.ID != 0 {
		return nil, errors.New("email already registered")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	newAdmin := &Admin{
		Name:        input.Name,
		PhoneNumber: input.Phone,
		Email:       input.Email,
		Password:    string(hash),
		Role:        input.Role,
	}

	if err := s.repo.Create(newAdmin); err != nil {
		return nil, err
	}
	return newAdmin, nil
}

func (s *service) GetAllAdmins() ([]Admin, error) {
	return s.repo.GetAll()
}
