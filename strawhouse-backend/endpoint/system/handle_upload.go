package system

import (
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/strawstacks/strawhouse/strawhouse-backend/type/payload"
	"github.com/strawstacks/strawhouse/strawhouse-backend/type/response"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/eventfeed"
	"github.com/strawstacks/strawhouse/strawhouse-proto/pb"
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

	// * Call upload service
	path, attribute, sum, encoded, er := r.File.Upload(token, fileHeader.Filename, directory, file)
	if er != nil {
		return er
	}

	// * Save file
	absolutePath := r.Filepath.AbsPath(*path)
	if err := c.SaveFile(fileHeader, absolutePath); err != nil {
		return gut.Err(false, "unable to save file", err)
	}

	// * Construct file flag
	if er := r.Fileflag.SumSet(*path, sum); er != nil {
		return er
	}

	// * Construct file flag
	if er := r.Fileflag.CorruptedInit(*path); er != nil {
		return er
	}

	// * Fire event feed
	r.EventFeed.Fire(eventfeed.FeedTypeUpload, *path, &pb.UploadFeedResponse{
		Name:      fileHeader.Filename,
		Directory: directory,
		Hash:      *encoded,
		Attr:      *attribute,
	})

	return c.JSON(response.Success(&payload.UploadResponse{
		Path: path,
		Hash: encoded,
		Size: gut.Ptr(fileHeader.Size),
	}))
}
