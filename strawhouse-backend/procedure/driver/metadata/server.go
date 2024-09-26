package metadata

import (
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/pogreb"
	"github.com/strawstacks/strawhouse/strawhouse-proto/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedDriverMetadataServer
	Pogreb *pogreb.Pogreb
}

func Init(registrar *grpc.Server, pogreb *pogreb.Pogreb) {
	server := &Server{
		Pogreb: pogreb,
	}

	pb.RegisterDriverMetadataServer(registrar, server)
}
