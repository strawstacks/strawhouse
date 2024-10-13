package feed

import (
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/eventfeed"
	"github.com/strawstacks/strawhouse/strawhouse-proto/pb"
)

func (r *Server) Upload(req *pb.UploadFeedRequest, stream pb.DriverFeed_UploadServer) error {
	// Bind the event
	id := r.EventFeed.Bind(eventfeed.FeedTypeUpload, req.Directory, func(resp any) {
		_ = stream.Send(resp.(*pb.UploadFeedResponse))
	})
	defer r.EventFeed.Unbind(eventfeed.FeedTypeUpload, req.Directory, id)

	<-stream.Context().Done()

	return nil
}
