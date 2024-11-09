package strawhouse

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/bsthun/gut"
	"github.com/strawst/strawhouse-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const GrpcMaxMessageSize = 1024 * 1024 * 1024 * 1024 * 1024 // 1 PB

type Clienter interface {
	Close() error
	DirectoryList(directory string) (*pb.DirectoryListResponse, error)
	TransferUpload(name string, directory string, content []byte, attribute []byte) error
	TransferGet(path string) (*pb.DownloadResponse, error)
	FeedUpload(directory string, callback func(resp *pb.UploadFeedResponse, err error)) (*FeedUploadSession, error)
}

type Client struct {
	Grpc                 *grpc.ClientConn
	driverMetadataClient pb.DriverMetadataClient
	driverTransferClient pb.DriverTransferClient
	driverFeedClient     pb.DriverFeedClient
}

func NewClient(option *Option) *Client {
	// * Construct credentials
	var cred credentials.TransportCredentials
	if option.Secure {
		roots, err := x509.SystemCertPool()
		if err != nil {
			gut.Fatal("failed to read system cert pool", err)
		}
		cred = credentials.NewTLS(&tls.Config{
			RootCAs:    roots,
			NextProtos: []string{"h2"},
		})
	} else {
		cred = insecure.NewCredentials()
	}

	// * Construct gRPC client
	gr, err := grpc.NewClient(
		option.Server,
		grpc.WithTransportCredentials(cred),
		grpc.WithUnaryInterceptor(UnaryInterceptor(option.Key)),
		grpc.WithStreamInterceptor(StreamInterceptor(option.Key)),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(GrpcMaxMessageSize),
			grpc.MaxCallSendMsgSize(GrpcMaxMessageSize),
		),
	)
	if err != nil {
		gut.Fatal("grpc failure", err)
	}

	driverMetadataClient := pb.NewDriverMetadataClient(gr)
	driverTransferClient := pb.NewDriverTransferClient(gr)
	driverFeedClient := pb.NewDriverFeedClient(gr)

	return &Client{
		Grpc:                 gr,
		driverMetadataClient: driverMetadataClient,
		driverTransferClient: driverTransferClient,
		driverFeedClient:     driverFeedClient,
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

func StreamInterceptor(key string) grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		md := metadata.New(map[string]string{"authorization": key})
		ctx = metadata.NewOutgoingContext(ctx, md)
		return streamer(ctx, desc, cc, method, opts...)
	}
}

func (r *Client) Close() error {
	return r.Grpc.Close()
}
