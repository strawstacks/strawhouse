package strawhouse

import (
	"github.com/strawstacks/strawhouse-go/pb"
	"golang.org/x/net/context"
)

func (r *Client) DirectoryList(directory string) (*pb.DirectoryListResponse, error) {
	resp, err := r.driverMetadataClient.DirectoryList(context.TODO(), &pb.DirectoryListRequest{
		Directory: directory,
	})
	return resp, err
}
