package grpc

import (
	"context"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func NewInterceptor(config *config.Config) *Interceptor {
	return &Interceptor{
		Config: config,
	}
}

type Interceptor struct {
	Config *config.Config
}

func (r *Interceptor) TokenAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Extract the metadata from the context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	// Extract the token from metadata
	key := md["authorization"]
	if len(key) == 0 {
		return nil, status.Error(codes.Unauthenticated, "authorization token not provided")
	}

	// Validate the token
	if key[0] == *r.Config.Key {
		return handler(ctx, req)
	}

	// Unauthorized
	return nil, status.Error(codes.Unauthenticated, "invalid token")
}
