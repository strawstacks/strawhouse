package get

import (
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/strawst/strawhouse-go"
	"mime"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func (r *Handler) Get(c *fiber.Ctx) error {
	// * Construct variables
	path := c.Path()
	path = filepath.Clean(path)
	path, err := url.PathUnescape(path)
	if err != nil {
		return gut.Err(false, "unable to decode path", err)
	}
	token := c.Query("t")

	// * Verify the file
	// TODO: Implement attribute usage from token
	if _, er := r.Signature.VerifyInt(strawhouse.SignatureActionGet, path, token); er != nil {
		return er
	}

	// * Construct absolute path
	absolutePath := r.Filepath.AbsPath(path)

	// * Check if path is a directory
	fileInfo, err := os.Stat(absolutePath)
	if err != nil || fileInfo.IsDir() {
		return gut.Err(false, "file not found")
	}

	// * Detect content type from file extension
	contentType := mime.TypeByExtension(filepath.Ext(absolutePath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// * Serve the file with the correct content type
	c.Set(fiber.HeaderContentType, contentType)
	c.Set(fiber.HeaderContentLength, strconv.FormatInt(fileInfo.Size(), 10))

	// * Stream the file
	writer := c.Response().BodyWriter()
	if err := r.File.Get(path, writer); err != nil {
		return err
	}

	return nil
}
