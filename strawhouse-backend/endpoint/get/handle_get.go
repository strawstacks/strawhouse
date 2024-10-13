package get

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/strawstacks/strawhouse/strawhouse-driver"
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
	fpath := filepath.Join(*r.Config.DataRoot, path)
	token := c.Query("t")
	attr := c.Query("a")

	// * Decode attributes
	r.Signature.ReplaceClean(&attr)
	attrBytes, err := base64.StdEncoding.DecodeString(attr)

	// * Verify the file
	if er := r.Signature.VerifyInt(strawhouse.SignatureActionGet, path, attrBytes, token); er != nil {
		return er
	}

	// * Check if path is a directory
	fileInfo, err := os.Stat(fpath)
	if err != nil || fileInfo.IsDir() {
		return gut.Err(false, "file not found")
	}

	// * Detect content type from file extension
	contentType := mime.TypeByExtension(filepath.Ext(fpath))
	if contentType == "" {
		contentType = "text/plain"
	}

	// * Serve the file with the correct content type
	c.Set(fiber.HeaderContentType, contentType)
	c.Set(fiber.HeaderContentLength, strconv.FormatInt(fileInfo.Size(), 10))

	// * Check file corruption
	if er := r.Fileflag.Corrupted(path); er != nil {
		return er
	}

	// * Check signed checksum
	sum, er := r.Fileflag.SumGet(path)
	if er != nil {
		return er
	}

	// * Check file path
	val, err := r.Pogreb.Sum.Get(sum)
	if err != nil || val == nil {
		return gut.Err(false, "file record not found")
	}
	if !bytes.Equal(val, []byte(path)) {
		return gut.Err(false, "file path mismatch")
	}

	// * Open the file
	file, err := os.Open(fpath)
	if err != nil {
		return gut.Err(false, "unable to open file", err)
	}
	defer func() {
		_ = file.Close()
	}()

	// * Initialize hash
	hash := sha256.New()
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			break
		}
		if n > 0 {
			hash.Write(buffer[:n])
			_, err := c.Write(buffer[:n])
			if err != nil {
				return gut.Err(false, err.Error())
			}
		}
	}

	// * Compare content hash and xattr hash
	if !bytes.Equal(hash.Sum(nil), sum) {
		if er := r.Fileflag.CorruptedSet(path, true); er != nil {
			return er
		}
		return gut.Err(false, "invalid file hash")
	}

	return nil
}
