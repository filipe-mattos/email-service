package campaign

import (
	"email-service/internal/contract"
	internal_errors "email-service/internal/internal-errors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaignDto) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internal_errors.ErrInternal
	}
	return campaign.ID, nil
}
