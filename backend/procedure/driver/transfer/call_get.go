package transfer

import (
	"bytes"
	"context"
	"github.com/strawst/strawhouse-go/pb"
	"path/filepath"
)

func (r *Server) FileDownloadPath(ctx context.Context, req *pb.DownloadPathRequest) (*pb.DownloadResponse, error) {
	// * Construct byte buffer
	buffer := new(bytes.Buffer)

	// * Call get service
	if er := r.File.Get(req.Path, buffer); er != nil {
		return nil, er
	}

	return &pb.DownloadResponse{
		Directory: filepath.Dir(req.Path),
		Name:      filepath.Base(req.Path),
		Content:   buffer.Bytes(),
	}, nil
}
