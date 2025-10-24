package admin

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetDashboardData() (map[string]interface{}, error) {
	stats, err := s.repo.GetStats()
	if err != nil {
		return nil, err
	}

	trend, err := s.repo.GetFinanceTrend()
	if err != nil {
		return nil, err
	}

	schedules, err := s.repo.GetSchedule()
	if err != nil {
		return nil, err
	}

	admins, err := s.repo.GetAdmins()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"stats":        stats,
		"financeTrend": trend,
		"schedule":     schedules,
		"admins":       admins,
	}, nil
}
