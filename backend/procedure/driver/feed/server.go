package feed

import (
	"github.com/strawstacks/strawhouse-go/pb"
	"google.golang.org/grpc"
	"strawhouse-backend/util/eventfeed"
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
