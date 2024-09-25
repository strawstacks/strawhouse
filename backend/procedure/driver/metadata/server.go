package metadata

import (
	"github.com/strawstacks/strawhouse/backend/common/pogreb"
	"google.golang.org/grpc"
	"proto/pb"
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
