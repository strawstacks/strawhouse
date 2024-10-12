package system

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/xattr"
	"github.com/strawstacks/strawhouse/strawhouse-backend/type/payload"
	"github.com/strawstacks/strawhouse/strawhouse-backend/type/response"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/signature"
	"github.com/strawstacks/strawhouse/strawhouse-driver"
	"os"
	"path/filepath"
)

func (r *Handler) Upload(c *fiber.Ctx) error {
	// * Parse body
	token := c.FormValue("token")
	destination := c.FormValue("destination")
	attribute := c.FormValue("attribute")
	fileHeader, err := c.FormFile("file")
	if token == "" || destination == "" {
		return gut.Err(false, "missing token, destination or attribute", nil)
	}
	if err != nil {
		return gut.Err(false, "unable to parse file", err)
	}
	var attrib []byte
	if attribute != "" {
		signature.ReplaceChar(&attribute, '+', '*')
		attrib, err = base64.StdEncoding.DecodeString(attribute)
		if err != nil {
			return gut.Err(false, "unable to decode attribute", err)
		}
	}

	// * Check file name
	fileHeader.Filename = r.Name.BaseName(fileHeader.Filename)
	if len(fileHeader.Filename) == 0 {
		return gut.Err(false, "invalid filename", nil)
	}

	// * Construct path
	directory := filepath.Clean(filepath.Join(*r.Config.DataRoot, destination))
	relativePath := filepath.Clean(filepath.Join(destination, fileHeader.Filename))
	absolutePath := filepath.Clean(filepath.Join(*r.Config.DataRoot, relativePath))

	// * Check token
	if err := r.Signature.Verify(strawhouse.SignatureActionUpload, relativePath, attrib, token); err != nil {
		return err
	}

	// * Ensure directory
	if err := os.MkdirAll(directory, 0700); err != nil {
		return gut.Err(false, "unable to create directory", err)
	}

	// * Open file
	file, err := fileHeader.Open()
	defer file.Close()
	if err != nil {
		return gut.Err(false, "unable to open file", err)
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
	if val, err := r.Pogreb.Sum.Get(sum); err != nil {
		return gut.Err(false, "unable to check hash", err)
	} else {
		if val != nil {
			pathVal := string(val)
			if pathVal != relativePath {
				return gut.Err(false, "file already exists in other path", nil)
			}
			if _, err := os.Stat(absolutePath); err == nil {
				flag, err := xattr.Get(absolutePath, "user.sh.flag")
				if err != nil {
					return gut.Err(false, "unable to get file flag attributes", err)
				}
				if flag[0]&0b00001000 == 0 {
					return gut.Err(false, "file is already exist", nil)
				}
			}
		}
	}

	// * Save file
	if err := c.SaveFile(fileHeader, absolutePath); err != nil {
		return gut.Err(false, "unable to save file", err)
	}

	// * Sign hash
	hash = r.Signature.GetHash()
	hash.Write(sum)
	signedSum := hash.Sum(nil)
	r.Signature.PutHash(hash)

	// Construct file attribute
	flagBytes := make([]byte, 4)

	// * Set file attributes
	err = xattr.Set(absolutePath, "user.sh.sum", sum)
	if err != nil {
		return gut.Err(false, "unable to set file sum attributes", err)
	}
	err = xattr.Set(absolutePath, "user.sh.flag", flagBytes)
	if err != nil {
		return gut.Err(false, "unable to set file flag attributes", err)
	}
	err = xattr.Set(absolutePath, "user.sh.sum.signed", signedSum)
	if err != nil {
		return gut.Err(false, "unable to set file signature attributes", err)
	}

	// * Save hash
	if err := r.Pogreb.Sum.Put(sum, []byte(relativePath)); err != nil {
		return gut.Err(false, "unable to save hash", err)
	}

	// * Save log
	size := r.Pogreb.Log.Count()
	sizeBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(sizeBytes, size)
	if err := r.Pogreb.Log.Put(sizeBytes, sum); err != nil {
		return gut.Err(false, "unable to save log", err)
	}

	// * Encode base64 hash
	encodedSum := base64.StdEncoding.EncodeToString(sum)

	return c.JSON(response.Success(&payload.UploadResponse{
		Path: gut.Ptr(relativePath),
		Hash: &encodedSum,
		Size: gut.Ptr(fileHeader.Size),
	}))
}
