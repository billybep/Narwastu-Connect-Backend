package sermon

import "time"

type Service interface {
	GetMonthlyArticles() ([]Sermon, error)
	GetByID(id uint) (*Sermon, error)
	GetWeeklyArticles() ([]Sermon, error)
	GetYearlyArticles(year int) ([]Sermon, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetWeeklyArticles() ([]Sermon, error) {
	return s.repo.FindLatestArticles(8) // ambil max 8
}

func (s *service) GetYearlyArticles(year int) ([]Sermon, error) {
	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)
	return s.repo.FindArticlesByDateRange(start, end)
}

func (s *service) GetMonthlyArticles() ([]Sermon, error) {
	return s.repo.GetMonthlyArticles()
}

func (s *service) GetByID(id uint) (*Sermon, error) {
	return s.repo.GetByID(id)
}
