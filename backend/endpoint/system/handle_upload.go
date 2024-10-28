package system

import (
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/strawstacks/strawhouse-go"
	"github.com/strawstacks/strawhouse-go/pb"
	"path/filepath"
	"strawhouse-backend/type/payload"
	"strawhouse-backend/type/response"
)

// Upload TODO: Implement body streaming
func (r *Handler) Upload(c *fiber.Ctx) error {
	// * Parse body
	token := c.FormValue("token")
	directory := c.FormValue("directory")
	fileHeader, err := c.FormFile("file")
	if token == "" || directory == "" {
		return gut.Err(false, "missing token or destination", nil)
	}
	if err != nil {
		return gut.Err(false, "unable to parse file", err)
	}

	// * Open file
	file, err := fileHeader.Open()
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		return gut.Err(false, "unable to open file", err)
	}

	// * Check token
	path := filepath.Clean(filepath.Join(directory, fileHeader.Filename))
	attribute, er := r.Signature.VerifyInt(strawhouse.SignatureActionUpload, path, token)
	if er != nil {
		return er
	}

	// * Call upload service
	_, sum, encoded, er := r.File.Upload(fileHeader.Filename, directory, file)
	if er != nil {
		return er
	}

	// * Save file
	absolutePath := r.Filepath.AbsPath(path)
	if err := c.SaveFile(fileHeader, absolutePath); err != nil {
		return gut.Err(false, "unable to save file", err)
	}

	// * Construct file flag
	if er := r.Fileflag.SumSet(path, sum); er != nil {
		return er
	}

	// * Construct file flag
	if er := r.Fileflag.CorruptedInit(path); er != nil {
		return er
	}

	// * Fire event feed
	r.EventFeed.Fire(strawhouse.FeedTypeUpload, path, &pb.UploadFeedResponse{
		Name:      fileHeader.Filename,
		Directory: directory,
		Hash:      *encoded,
		Attr:      attribute,
	})

	return c.JSON(response.Success(&payload.UploadResponse{
		Path: &path,
		Hash: encoded,
		Size: gut.Ptr(fileHeader.Size),
	}))
}
