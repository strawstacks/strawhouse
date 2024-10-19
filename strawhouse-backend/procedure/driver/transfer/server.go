package transfer

import (
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/pogreb"
	"github.com/strawstacks/strawhouse/strawhouse-backend/service/file"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/eventfeed"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/fileflag"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/filepath"
	"github.com/strawstacks/strawhouse/strawhouse-proto/pb"
	"google.golang.org/grpc"
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
