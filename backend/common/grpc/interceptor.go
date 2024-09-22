package grpc

import (
	"backend/common/config"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func NewInterceptor(config *config.Config) *Interceptor {
	return &Interceptor{
		config: config,
	}
}

type Interceptor struct {
	config *config.Config
}

func (r *Interceptor) TokenAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Extract the metadata from the context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	// Extract the token from metadata
	tokens := md["authorization"]
	if len(tokens) == 0 {
		return nil, status.Error(codes.Unauthenticated, "authorization token not provided")
	}

	token := tokens[0]

	// Validate the token
	for _, client := range r.config.Clients {
		if token == *client.Key {
			return handler(ctx, req)
		}
	}

	// Unauthorized
	return nil, status.Error(codes.Unauthenticated, "invalid token")
}
