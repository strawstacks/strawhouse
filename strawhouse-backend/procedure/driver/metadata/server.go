package metadata

import (
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/pogreb"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/eventfeed"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/filepath"
	"github.com/strawstacks/strawhouse/strawhouse-proto/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedDriverMetadataServer
	Config    *config.Config
	Pogreb    *pogreb.Pogreb
	Filepath  *filepath.Filepath
	EventFeed *eventfeed.EventFeed
}

func Register(registrar *grpc.Server, config *config.Config, pogreb *pogreb.Pogreb, filepath *filepath.Filepath, eventfeed *eventfeed.EventFeed) {
	server := &Server{
		Config:    config,
		Pogreb:    pogreb,
		Filepath:  filepath,
		EventFeed: eventfeed,
	}

	pb.RegisterDriverMetadataServer(registrar, server)
}
