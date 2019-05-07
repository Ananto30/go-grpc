package usecase

import (
	"github.com/Ananto30/go-grpc/campaign"
	"github.com/Ananto30/go-grpc/model"
)

type campaignUsecase struct {
	cmpRepo campaign.Repository
}

// NewCampaignUsecase will create new an campaignUsecase object representation of campaign.Usecase interface
func NewCampaignUsecase(cr campaign.Repository) campaign.Usecase {
	return &campaignUsecase{cmpRepo: cr}
}

func (cu *campaignUsecase) GetByID(id string) (*model.Campaign, error) {
	cmp, err := cu.cmpRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return cmp, nil
}

func (cu *campaignUsecase) Store(c campaign.Campaign) (*model.Campaign, error) {
	cmp, err := cu.cmpRepo.Store(*c.ToModel())
	if err != nil {
		return nil, err
	}
	return cmp, nil
}
