package registry

import (
	"go.uber.org/dig"
	campaign_rpc "magic.pathao.com/nightwing/fortress/campaign/delivery/rpc"
	campaign_repository "magic.pathao.com/nightwing/fortress/campaign/repository"
	campaign_usecase "magic.pathao.com/nightwing/fortress/campaign/usecase"
	"magic.pathao.com/nightwing/fortress/server"
)

// BuildContainer ...
func BuildContainer() *dig.Container {
	ctn := dig.New()
	ctn.Provide(server.NewServer)
	ctn.Provide(campaign_repository.NewMGORepository)
	ctn.Provide(campaign_usecase.NewCampaignUsecase)
	ctn.Provide(campaign_rpc.NewCampaignServer)
	return ctn
}
