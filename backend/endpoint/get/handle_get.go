package get

import (
	"backend/util/signature"
	uu "github.com/bsthun/goutils"
	"github.com/gofiber/fiber/v2"
	"mime"
	"os"
	"path/filepath"
	"strconv"
)

func (r *Handler) Get(c *fiber.Ctx) error {
	// * Construct variables
	root := *r.Config.DataRoot
	path := c.Path()
	token := c.Query("t")

	// * Concatenate path
	fpath := filepath.Join(root, path)

	// * Check if path is a directory
	fileInfo, err := os.Stat(fpath)
	if err != nil || fileInfo.IsDir() {
		return fiber.ErrNotFound
	}

	// * Verify the file
	if err := signature.Verify(path, token, *r.Config.Clients[0].Key); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(uu.Err(false, err.Error()))
	}

	// * Open the file
	file, err := os.Open(fpath)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(uu.Err(false, err.Error()))
	}
	defer file.Close()

	// * Skip first 64 bytes of the file
	_, err = file.Seek(64, 0) // Skip first 64 bytes
	if err != nil {
		return fiber.ErrInternalServerError
	}

	// * Detect content type from file extension
	contentType := mime.TypeByExtension(filepath.Ext(fpath))
	if contentType == "" {
		contentType = "text/plain"
	}

	// * Serve the file with the correct content type, skipping the first 64 bytes
	c.Set(fiber.HeaderContentType, contentType)
	c.Set(fiber.HeaderContentLength, strconv.FormatInt(fileInfo.Size()-64, 10))

	// Stream file from 64th byte onward
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			break
		}
		if n > 0 {
			_, err := c.Write(buffer[:n])
			if err != nil {
				return uu.Err(false, err.Error())
			}
		}
	}

	return nil
}
