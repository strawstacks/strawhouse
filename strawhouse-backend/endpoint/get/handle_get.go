package get

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/xattr"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/signature"
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
	signature.ReplaceChar(&attr, '*', '+')
	attrBytes, err := base64.StdEncoding.DecodeString(attr)

	// * Verify the file
	if err := r.Signature.Verify(strawhouse.SignatureActionGet, path, attrBytes, token); err != nil {
		return err
	}

	// * Check if path is a directory
	fileInfo, err := os.Stat(fpath)
	if err != nil || fileInfo.IsDir() {
		return gut.Err(false, "File not found")
	}

	// * Open the file
	file, err := os.Open(fpath)
	if err != nil {
		return gut.Err(false, "unable to open file", err)
	}
	defer file.Close()

	// * Detect content type from file extension
	contentType := mime.TypeByExtension(filepath.Ext(fpath))
	if contentType == "" {
		contentType = "text/plain"
	}

	// * Serve the file with the correct content type, skipping the first 64 bytes
	c.Set(fiber.HeaderContentType, contentType)
	c.Set(fiber.HeaderContentLength, strconv.FormatInt(fileInfo.Size(), 10))

	// * Check file flag
	flag, err := xattr.Get(fpath, "user.sh.flag")
	if err != nil {
		return gut.Err(false, "unable to get file flag attributes", err)
	}
	if flag[0]&0b00001000 != 0 {
		return gut.Err(false, "file corrupted")
	}

	// * Check file attribute
	sum, err := xattr.Get(fpath, "user.sh.sum")
	if err != nil {
		return gut.Err(false, "unable to get file sum attributes", err)
	}
	signedSum, err := xattr.Get(fpath, "user.sh.sum.signed")
	if err != nil {
		return gut.Err(false, "unable to set file signature attributes", err)
	}

	// * Check file path
	val, err := r.Pogreb.Sum.Get(sum)
	if err != nil || val == nil {
		return gut.Err(false, "file record not found")
	}
	if !bytes.Equal(val, []byte(path)) {
		return gut.Err(false, "file path mismatch")
	}

	// * Validate the file
	hash := r.Signature.GetHash()
	hash.Write(sum)
	if !bytes.Equal(hash.Sum(nil), signedSum) {
		return gut.Err(false, "invalid file signature")
	}
	r.Signature.PutHash(hash)

	// * Initialize the hash
	hash = sha256.New()
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

	// * Check the hash
	if !bytes.Equal(hash.Sum(nil), sum) {
		flag[0] |= 0b00001000
		_ = xattr.Set(fpath, "user.sh.flag", flag)
		return gut.Err(false, "invalid file hash")
	}

	return nil
}
