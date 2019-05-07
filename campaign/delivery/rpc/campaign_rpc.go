package rpc

import (
	"context"

	"github.com/Ananto30/go-grpc/campaign"
	"github.com/Ananto30/go-grpc/helpers"
	pb "github.com/Ananto30/go-grpc/proto/fortress"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type campaignServer struct {
	cmpUsecase campaign.Usecase
}

// NewCampaignServer ...
func NewCampaignServer(gserver *grpc.Server, cmpUsecase campaign.Usecase) *campaignServer {

	server := &campaignServer{
		cmpUsecase: cmpUsecase,
	}

	pb.RegisterFortressServer(gserver, server)
	reflection.Register(gserver)
	return server
}

// GetCampaign returns the campaign identified by ID
func (s *campaignServer) GetCampaign(_ context.Context, r *pb.ReqGetCampaign) (*pb.RespGetCampaign, error) {
	// cmp, err := s.data.GetCampaignByID(r.GetId())
	// if err != nil {
	// 	// TODO:: Implement Common Error
	// 	panic("unimplemented")
	// }

	// resp := &pb.RespGetCampaign{
	// 	// Id:   cmp.ID,
	// 	// Name: cmp.Name,
	// }

	return nil, nil
}

// CreateCampaign creates a campaign and returns it
func (s *campaignServer) CreateCampaign(_ context.Context, r *pb.ReqCreateCampaign) (*pb.RespCreateCampaign, error) {
	c, errV := formatCreateCampaign(r)
	if errV != nil {
		resp := &pb.RespCreateCampaign{
			Error: errV.Error(),
		}
		return resp, nil
	}
	cmp, err := s.cmpUsecase.Store(*c)
	if err != nil {
		resp := &pb.RespCreateCampaign{
			Error: helpers.FormatError(err),
		}
		return resp, nil
	}

	resp := &pb.RespCreateCampaign{
		Campaign: formatCampaign(campaign.FormatCampaign(cmp)),
	}
	return resp, nil
}
