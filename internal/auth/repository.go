package auth

import "app/pkg/repository"

type UserRepository struct{}

func NewUserRepository() *UserRepository { return &UserRepository{} }

func (r *UserRepository) FindByProvider(pid, provider string) (*User, error) {
	var u User
	if err := repository.DB.Where("provider = ? AND provider_id = ?", provider, pid).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Create(u *User) error {
	return repository.DB.Create(u).Error
}

func (r *UserRepository) Save(u *User) error {
	return repository.DB.Save(u).Error
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
	var u User
	if err := repository.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) DeleteUser(userID uint) error {
	// Cascade otomatis handle VerseLikes karena sudah ada OnDelete:CASCADE
	return repository.DB.Delete(&User{}, userID).Error
}
