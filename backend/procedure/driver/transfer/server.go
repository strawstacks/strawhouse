package transfer

import (
	"github.com/strawst/strawhouse-go/pb"
	"google.golang.org/grpc"
	"strawhouse-backend/common/config"
	"strawhouse-backend/common/pogreb"
	"strawhouse-backend/service/file"
	"strawhouse-backend/util/eventfeed"
	"strawhouse-backend/util/fileflag"
	"strawhouse-backend/util/filepath"
)

type Server struct {
	pb.UnimplementedDriverTransferServer
	Config    *config.Config
	Pogreb    *pogreb.Pogreb
	File      *file.Service
	Filepath  *filepath.Filepath
	Fileflag  *fileflag.Fileflag
	EventFeed *eventfeed.EventFeed
}

func Register(registrar *grpc.Server, config *config.Config, pogreb *pogreb.Pogreb, file *file.Service, filepath *filepath.Filepath, fileflag *fileflag.Fileflag, eventfeed *eventfeed.EventFeed) {
	server := &Server{
		Config:    config,
		Pogreb:    pogreb,
		File:      file,
		Filepath:  filepath,
		Fileflag:  fileflag,
		EventFeed: eventfeed,
	}

	pb.RegisterDriverTransferServer(registrar, server)
}
