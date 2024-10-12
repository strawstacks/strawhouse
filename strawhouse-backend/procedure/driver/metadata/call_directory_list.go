package metadata

import (
	"context"
	"github.com/strawstacks/strawhouse/strawhouse-proto/pb"
)

func (r *Server) DirectoryList(ctx context.Context, req *pb.DirectoryListRequest) (*pb.DirectoryListResponse, error) {
	return &pb.DirectoryListResponse{}, nil
}
