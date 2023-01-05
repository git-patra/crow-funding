package campaign

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo,
	}
}

func (s *service) FindCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		return s.repository.FindByUserID(userID)
	}

	return s.repository.FindAll()
}
