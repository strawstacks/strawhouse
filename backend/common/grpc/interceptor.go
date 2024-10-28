package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strawhouse-backend/common/config"
)

func NewInterceptor(config *config.Config) *Interceptor {
	return &Interceptor{
		Config: config,
	}
}

type Interceptor struct {
	Config *config.Config
}

func (r *Interceptor) AuthorizationUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
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

func (r *Interceptor) AuthorizationStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// Extract the metadata from the stream context
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return status.Error(codes.Unauthenticated, "missing metadata")
	}

	// Extract the token from metadata
	key := md["authorization"]
	if len(key) == 0 {
		return status.Error(codes.Unauthenticated, "authorization token not provided")
	}

	// Validate the token
	if key[0] == *r.Config.Key {
		return handler(srv, ss)
	}

	// Unauthorized
	return status.Error(codes.Unauthenticated, "invalid token")
}
