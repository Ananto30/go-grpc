package server

import (
	"log"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	campaign_rpc "magic.pathao.com/nightwing/fortress/campaign/delivery/rpc"
	campaign_repository "magic.pathao.com/nightwing/fortress/campaign/repository"
	campaign_usecase "magic.pathao.com/nightwing/fortress/campaign/usecase"
	"magic.pathao.com/nightwing/fortress/conn"
)

// Server represets the fortress rpc server
type Server struct {
	*grpc.Server
}

var logger = logrus.NewEntry(logrus.New())
var unaryInterceptors = []grpc.UnaryServerInterceptor{
	grpc_logrus.UnaryServerInterceptor(logger),
}

// NewServer returns a new grpc server instance
func NewServer() *Server {

	srvr := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)),
	)

	s := &Server{Server: srvr}
	return s
}

func (s *Server) registerCampaign() {
	repo, err := campaign_repository.NewMGORepository(conn.DefaultDB().Session)
	if err != nil {
		log.Fatalf("Failed to connect to MGO: %v", err)
	}
	uc := campaign_usecase.NewCampaignUsecase(repo)
	campaign_rpc.NewCampaignServer(s.Server, uc)
}
