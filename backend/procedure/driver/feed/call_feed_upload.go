package feed

import (
	"github.com/strawstacks/strawhouse-go"
	"github.com/strawstacks/strawhouse-go/pb"
)

func (r *Server) Upload(req *pb.UploadFeedRequest, stream pb.DriverFeed_UploadServer) error {
	// Bind the event
	id := r.EventFeed.Bind(strawhouse.FeedTypeUpload, req.Directory, func(resp any) {
		_ = stream.Send(resp.(*pb.UploadFeedResponse))
	})
	defer r.EventFeed.Unbind(strawhouse.FeedTypeUpload, req.Directory, id)

	<-stream.Context().Done()

	return nil
}
