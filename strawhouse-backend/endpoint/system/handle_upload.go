package system

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/strawstacks/strawhouse/strawhouse-backend/type/payload"
	"github.com/strawstacks/strawhouse/strawhouse-backend/type/response"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/eventfeed"
	"github.com/strawstacks/strawhouse/strawhouse-driver"
	"github.com/strawstacks/strawhouse/strawhouse-proto/pb"
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
		r.Signature.ReplaceUnclean(&attribute)
		attrib, err = base64.StdEncoding.DecodeString(attribute)
		if err != nil {
			return gut.Err(false, "unable to decode attribute", err)
		}
	}

	// * Check file name
	fileHeader.Filename = r.Filepath.BaseName(fileHeader.Filename)
	if len(fileHeader.Filename) == 0 {
		return gut.Err(false, "invalid filename", nil)
	}

	// * Construct path
	directory := filepath.Clean(filepath.Join(*r.Config.DataRoot, destination))
	relativePath := filepath.Clean(filepath.Join(destination, fileHeader.Filename))
	absolutePath := filepath.Clean(filepath.Join(*r.Config.DataRoot, relativePath))

	// * Check token
	if err := r.Signature.VerifyInt(strawhouse.SignatureActionUpload, relativePath, attrib, token); err != nil {
		return err
	}

	// * Ensure directory
	if err := os.MkdirAll(directory, 0700); err != nil {
		return gut.Err(false, "unable to create directory", err)
	}

	// * Open file
	file, err := fileHeader.Open()
	defer func() {
		_ = file.Close()
	}()
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
			// Check if file already exists in other path
			pathVal := string(val)
			if pathVal != relativePath {
				return gut.Err(false, "file already exists in other path", nil)
			}

			// If file exists and not corrupted, deny upload
			if _, err := os.Stat(absolutePath); err == nil {
				if er := r.Fileflag.Corrupted(relativePath); er == nil {
					return gut.Err(false, "file is already exist", nil)
				}
			}
		}
	}

	// * Save file
	if err := c.SaveFile(fileHeader, absolutePath); err != nil {
		return gut.Err(false, "unable to save file", err)
	}

	// Construct file flag
	if er := r.Fileflag.CorruptedInit(relativePath); er != nil {
		return er
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
	encodedSum = encodedSum[:len(encodedSum)-1]
	r.Signature.ReplaceClean(&encodedSum)

	// * Fire event feed
	r.EventFeed.Fire(eventfeed.FeedTypeUpload, relativePath, &pb.UploadFeedResponse{
		Name:      fileHeader.Filename,
		Directory: destination,
		Hash:      encodedSum,
		Attr:      attrib,
	})

	return c.JSON(response.Success(&payload.UploadResponse{
		Path: gut.Ptr(relativePath),
		Hash: &encodedSum,
		Size: gut.Ptr(fileHeader.Size),
	}))
}
