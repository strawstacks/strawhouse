package strawhouse

import (
	"context"
	"github.com/bsthun/gut"
	"github.com/strawstacks/strawhouse/strawhouse-proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type Clienter interface {
	Close() error
	DirectoryList(directory string) (*pb.DirectoryListResponse, error)
}

type Client struct {
	Grpc                 *grpc.ClientConn
	driverMetadataClient pb.DriverMetadataClient
	driverTransferClient pb.DriverTransferClient
}

func NewClient(key string, server string) *Client {
	gr, err := grpc.NewClient(
		server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(UnaryInterceptor(key)),
	)
	if err != nil {
		gut.Fatal("grpc failure", err)
	}

	driverMetadataClient := pb.NewDriverMetadataClient(gr)
	driverTransferClient := pb.NewDriverTransferClient(gr)

	return &Client{
		Grpc:                 gr,
		driverMetadataClient: driverMetadataClient,
		driverTransferClient: driverTransferClient,
	}
}

func UnaryInterceptor(key string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req any,
		reply any,
		cc *grpc.ClientConn,
		invoke grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		md := metadata.New(map[string]string{"authorization": key})
		ctx = metadata.NewOutgoingContext(ctx, md)
		return invoke(ctx, method, req, reply, cc, opts...)
	}
}

func (r *Client) Close() error {
	return r.Grpc.Close()
}
