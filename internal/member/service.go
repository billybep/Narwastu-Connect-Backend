package member

import (
	"time"
)

type Service interface {
	GetWeeklyBirthdays() ([]Member, error)
	CreateMember(member *Member) error
	GetAllMembers() ([]Member, error)
	GetMemberByID(id uint) (*Member, error)
	UpdateMember(id uint, data *Member) (*Member, error)
	DeleteMember(id uint) error

	//Upload Profile
	UpdateMemberAvatar(memberID uint, avatarURL string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAllMembers() ([]Member, error) {
	return s.repo.FindAllMembers()
}

func (s *service) GetMemberByID(id uint) (*Member, error) {
	return s.repo.FindMemberByID(id)
}

func (s *service) UpdateMember(id uint, data *Member) (*Member, error) {
	existing, err := s.repo.FindMemberByID(id)
	if err != nil {
		return nil, err
	}

	// Update field yang boleh diubah
	existing.FullName = data.FullName
	existing.FamilyName = data.FamilyName
	existing.PlaceOfBirth = data.PlaceOfBirth
	existing.DateOfBirth = data.DateOfBirth
	existing.Gender = data.Gender
	existing.Address = data.Address
	existing.Phone = data.Phone
	existing.Email = data.Email
	existing.BaptismDate = data.BaptismDate
	existing.SiteID = data.SiteID
	existing.MaritalStatus = data.MaritalStatus
	existing.PhotoURL = data.PhotoURL
	existing.IsActive = data.IsActive

	if err := s.repo.UpdateMember(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *service) GetWeeklyBirthdays() ([]Member, error) {
	now := time.Now()
	// range minggu: Minggu (Sunday) - Sabtu (Saturday)
	weekday := int(now.Weekday()) // 0=Sunday
	start := now.AddDate(0, 0, -weekday)
	end := start.AddDate(0, 0, 6)

	return s.repo.FindBirthdaysInRange(start, end)
}

func (s *service) CreateMember(member *Member) error {
	return s.repo.CreateMember(member)
}

func (s *service) DeleteMember(id uint) error {
	// Pastikan member ada dulu sebelum hapus
	_, err := s.repo.FindMemberByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteMember(id)
}

func (s *service) UpdateMemberAvatar(memberID uint, avatarURL string) error {
	return s.repo.UpdateMemberAvatar(memberID, avatarURL)
}
