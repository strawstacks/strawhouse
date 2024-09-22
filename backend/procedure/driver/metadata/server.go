package metadata

import (
	"google.golang.org/grpc"
	"proto/pb"
)

type Server struct {
	pb.UnimplementedDriverMetadataServer
}

func Init(registrar *grpc.Server) {
	server := &Server{}

	pb.RegisterDriverMetadataServer(registrar, server)
}
