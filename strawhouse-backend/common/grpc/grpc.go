package grpc

import (
	"context"
	"github.com/bsthun/gut"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"net"
)

func Init(lc fx.Lifecycle, config *config.Config) *grpc.Server {
	// * Initialize interceptor
	interceptor := NewInterceptor(config)

	// * Initialize gRPC server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.TokenAuthInterceptor),
	)

	// * Append lifecycle hook
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				lis, err := net.Listen(*config.ProtoListen[0], *config.ProtoListen[1])
				if err != nil {
					gut.Fatal("Unable to listen", err)
				}
				if err := grpcServer.Serve(lis); err != nil {
					gut.Fatal("Unable to serve", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
	})

	// * Return gRPC server
	return grpcServer
}
