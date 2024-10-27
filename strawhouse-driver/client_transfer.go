package strawhouse

import (
	"github.com/strawstacks/strawhouse/strawhouse-proto/pb"
	"golang.org/x/net/context"
)

func (r *Client) TransferUpload(name string, directory string, content []byte, attribute []byte) error {
	_, err := r.driverTransferClient.FileUpload(
		context.TODO(),
		&pb.UploadRequest{
			Name:      name,
			Directory: directory,
			Content:   content,
			Attribute: attribute,
		},
	)
	return err
}

func (r *Client) TransferGet(path string) (*pb.DownloadResponse, error) {
	resp, err := r.driverTransferClient.FileDownloadPath(
		context.TODO(),
		&pb.DownloadPathRequest{
			Path: path,
		},
	)

	return resp, err
}
