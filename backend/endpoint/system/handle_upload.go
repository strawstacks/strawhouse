package system

import (
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/strawstacks/strawhouse-go"
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
	_, _, encoded, er := r.File.Upload(fileHeader.Filename, directory, attribute, file)
	if er != nil {
		return er
	}

	return c.JSON(response.Success(&payload.UploadResponse{
		Path: &path,
		Hash: encoded,
		Size: gut.Ptr(fileHeader.Size),
	}))
}
