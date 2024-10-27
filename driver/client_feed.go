package strawhouse

import (
	"context"
	"github.com/strawstacks/strawhouse-go/pb"
	"google.golang.org/grpc"
)

type FeedUploadSession struct {
	client   grpc.ServerStreamingClient[pb.UploadFeedResponse]
	callback func(resp *pb.UploadFeedResponse, err error)
	closeCh  chan struct{}
}

func (r *FeedUploadSession) Close() {
	close(r.closeCh)
}

func (r *Client) FeedUpload(directory string, callback func(resp *pb.UploadFeedResponse, err error)) (*FeedUploadSession, error) {
	client, err := r.driverFeedClient.Upload(context.TODO(), &pb.UploadFeedRequest{
		Directory: directory,
	})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = client.CloseSend()
	}()

	// * Construct session
	session := &FeedUploadSession{
		client:   client,
		callback: callback,
		closeCh:  make(chan struct{}),
	}

	// * Start listening
	go func() {
		for {
			select {
			case <-session.closeCh:
				return
			default:
				resp, err := session.client.Recv()
				if err != nil {
					session.callback(nil, err)
					return
				}
				session.callback(resp, nil)
			}
		}
	}()

	return session, err
}
