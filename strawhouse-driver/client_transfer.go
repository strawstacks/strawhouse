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
