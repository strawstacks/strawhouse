package metadata

import (
	"context"
	"github.com/strawstacks/strawhouse-go/pb"
	"os"
	"path/filepath"
)

func (r *Server) DirectoryList(ctx context.Context, req *pb.DirectoryListRequest) (*pb.DirectoryListResponse, error) {
	dir := filepath.Join(*r.Config.DataRoot, req.Directory)

	files := make([]*pb.File, 0)
	directories := make([]*pb.Directory, 0)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if path != dir {
				directories = append(directories, &pb.Directory{
					Name: info.Name(),
					Path: r.Filepath.RelPath(path) + "/",
				})
			}
		} else {
			files = append(files, &pb.File{
				Name:      info.Name(),
				Directory: r.Filepath.RelPath(filepath.Dir(path)) + "/",
				Checksum:  "",
				Size:      info.Size(),
				Mtime:     info.ModTime().Unix(),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.DirectoryListResponse{
		Name:        req.Directory,
		Files:       files,
		Directories: directories,
	}, nil
}
