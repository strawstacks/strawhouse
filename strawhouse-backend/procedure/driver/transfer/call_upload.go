package transfer

import (
	"bytes"
	"context"
	"github.com/bsthun/gut"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/eventfeed"
	"github.com/strawstacks/strawhouse/strawhouse-proto/pb"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"os"
)

func (r *Server) Upload(ctx context.Context, req *pb.UploadRequest) (*emptypb.Empty, error) {
	// * Create io.Reader
	reader := bytes.NewReader(req.Content)
	readCloser := io.NopCloser(reader)

	// * Call upload service
	path, sum, encoded, er := r.File.Upload(req.Name, req.Directory, readCloser)
	if er != nil {
		return nil, er
	}

	// * Create file
	file, err := os.Create(*path)
	if err != nil {
		return nil, gut.Err(false, "unable to create file", err)
	}
	defer func() {
		_ = file.Close()
	}()

	// * Save file
	if _, err := io.Copy(file, reader); err != nil {
		return nil, gut.Err(false, "unable to save file", err)
	}

	// * Construct file flag
	if er := r.Fileflag.SumSet(*path, sum); er != nil {
		return nil, er
	}

	// * Construct file flag
	if er := r.Fileflag.CorruptedInit(*path); er != nil {
		return nil, er
	}

	// * Fire event feed
	r.EventFeed.Fire(eventfeed.FeedTypeUpload, *path, &pb.UploadFeedResponse{
		Name:      req.Name,
		Directory: req.Directory,
		Hash:      *encoded,
		Attr:      req.Attribute,
	})

	return new(emptypb.Empty), nil
}
