package organization

type Service interface {
	GetAll() ([]Organization, error)
	GetByCategory(category string) ([]Organization, error)
	Create(org Organization) error
	Update(id uint, org Organization) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]Organization, error) {
	return s.repo.FindAll()
}

func (s *service) GetByCategory(category string) ([]Organization, error) {
	return s.repo.FindByCategory(category)
}

func (s *service) Create(org Organization) error {
	return s.repo.Create(org)
}

func (s *service) Update(id uint, org Organization) error {
	return s.repo.Update(id, org)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}
