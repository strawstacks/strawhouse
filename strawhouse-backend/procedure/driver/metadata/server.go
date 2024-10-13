package metadata

import (
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/pogreb"
	"github.com/strawstacks/strawhouse/strawhouse-proto/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedDriverMetadataServer
	Config *config.Config
	Pogreb *pogreb.Pogreb
}

func Init(registrar *grpc.Server, config *config.Config, pogreb *pogreb.Pogreb) {
	server := &Server{
		Config: config,
		Pogreb: pogreb,
	}

	pb.RegisterDriverMetadataServer(registrar, server)
}
