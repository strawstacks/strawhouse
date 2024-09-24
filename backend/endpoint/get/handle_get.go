package get

import (
	"backend/type/enum"
	"backend/util/signature"
	"encoding/base64"
	uu "github.com/bsthun/goutils"
	"github.com/gofiber/fiber/v2"
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
		return uu.Err(false, "Unable to decode path", err)
	}
	fpath := filepath.Join(*r.Config.DataRoot, path)
	token := c.Query("t")
	attr := c.Query("a")

	// * Decode attributes
	signature.ReplaceChar(&attr, '*', '+')
	attrBytes, err := base64.StdEncoding.DecodeString(attr)

	// * Verify the file
	if err := r.Signature.Verify(enum.SignatureActionGet, path, attrBytes, token); err != nil {
		return err
	}

	// * Check if path is a directory
	fileInfo, err := os.Stat(fpath)
	if err != nil || fileInfo.IsDir() {
		return uu.Err(false, "File not found")
	}

	// * Open the file
	file, err := os.Open(fpath)
	if err != nil {
		return uu.Err(false, "Unable to open file", err)
	}
	defer file.Close()

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
