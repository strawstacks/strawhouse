package system

import (
	"backend/type/enum"
	"backend/type/payload"
	"backend/type/response"
	"backend/util/signature"
	"crypto/sha256"
	"encoding/base64"
	uu "github.com/bsthun/goutils"
	"github.com/gofiber/fiber/v2"
	"os"
	"path/filepath"
)

func (r *Handler) Upload(c *fiber.Ctx) error {
	// * Parse body
	token := c.FormValue("token")
	destination := c.FormValue("destination")
	attribute := c.FormValue("attribute")
	fileHeader, err := c.FormFile("file")
	if token == "" || destination == "" || attribute == "" {
		return uu.Err(false, "missing token, destination or attribute", nil)
	}
	if err != nil {
		return uu.Err(false, "unable to parse file", err)
	}
	signature.ReplaceChar(&attribute, '+', '*')
	attrib, err := base64.StdEncoding.DecodeString(attribute)
	if err != nil {
		return uu.Err(false, "unable to decode attribute", err)
	}

	// * Check file name
	fileHeader.Filename = r.Name.BaseName(fileHeader.Filename)
	if len(fileHeader.Filename) == 0 {
		return uu.Err(false, "invalid filename", nil)
	}

	// * Construct path
	directory := filepath.Clean(filepath.Join(*r.Config.DataRoot, destination))
	relativePath := filepath.Clean(filepath.Join(destination, fileHeader.Filename))
	absolutePath := filepath.Clean(filepath.Join(*r.Config.DataRoot, relativePath))

	// * Ensure directory
	if err := os.MkdirAll(directory, 0700); err != nil {
		return uu.Err(false, "unable to create directory", err)
	}

	// * Check token
	if err := r.Signature.Verify(enum.SignatureActionUpload, relativePath, attrib, token); err != nil {
		return err
	}

	// * Open file
	file, err := fileHeader.Open()
	if err != nil {
		return uu.Err(false, "unable to open file", err)
	}

	// * Calculate sha256 hash
	hash := sha256.New()
	fileBuffer := make([]byte, 1024)
	for {
		n, err := file.Read(fileBuffer)
		if err != nil {
			break
		}
		hash.Write(fileBuffer[:n])
	}
	sum := hash.Sum(nil)

	// * Check hash
	if result, err := r.Pogreb.Get(sum); err != nil {
		return uu.Err(false, "unable to check hash", err)
	} else {
		if result != nil {
			return uu.Err(false, "file already exists", nil)
		}
	}

	// * Save file
	if err := c.SaveFile(fileHeader, absolutePath); err != nil {
		return uu.Err(false, "unable to save file", err)
	}

	// * Encode base64 hash
	encodedSum := base64.StdEncoding.EncodeToString(sum)

	return c.JSON(response.Success(&payload.UploadResponse{
		Path: uu.Ptr(relativePath),
		Hash: &encodedSum,
		Size: uu.Ptr(fileHeader.Size),
	}))
}
