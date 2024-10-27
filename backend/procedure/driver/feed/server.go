package feed

import (
	"github.com/strawstacks/strawhouse-go/pb"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/eventfeed"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedDriverFeedServer
	EventFeed *eventfeed.EventFeed
}

func Register(registrar *grpc.Server, eventfeed *eventfeed.EventFeed) {
	server := &Server{
		EventFeed: eventfeed,
	}

	pb.RegisterDriverFeedServer(registrar, server)
}
