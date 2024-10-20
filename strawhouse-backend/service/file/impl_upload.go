package file

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"github.com/bsthun/gut"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func (r *Service) Upload(name string, directory string, file io.ReadCloser) (*string, []byte, *string, *gut.ErrorInstance) {
	// * Validate name
	if strings.HasPrefix(name, ".") {
		return nil, nil, nil, gut.Err(false, "invalid filename", nil)
	}
	if len(name) < 3 {
		return nil, nil, nil, gut.Err(false, "invalid filename", nil)
	}

	// * Normalize file name
	name = r.filepath.BaseName(name)

	// * Construct path
	relativeFilePath := filepath.Clean(filepath.Join(directory, name))
	absoluteFilePath := r.filepath.AbsPath(filepath.Clean(relativeFilePath))
	absoluteDirectoryPath := r.filepath.AbsPath(directory)

	// * Ensure directory
	if err := os.MkdirAll(absoluteDirectoryPath, 0700); err != nil {
		return nil, nil, nil, gut.Err(false, "unable to create directory", err)
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
	if val, err := r.pogreb.Sum.Get(sum); err != nil {
		return nil, nil, nil, gut.Err(false, "unable to check hash", err)
	} else {
		if val != nil {
			// Check if file already exists in other path
			pathVal := string(val)
			if pathVal != relativeFilePath {
				return nil, nil, nil, gut.Err(false, "file already exists in other path", nil)
			}

			// If file exists and not corrupted, deny upload
			if _, err := os.Stat(absoluteFilePath); err == nil {
				if er := r.fileflag.Corrupted(relativeFilePath); er == nil {
					return nil, nil, nil, gut.Err(false, "file already exist", nil)
				}
			}
		}
	}

	// * Save hash
	if err := r.pogreb.Sum.Put(sum, []byte("/"+relativeFilePath)); err != nil {
		return nil, nil, nil, gut.Err(false, "unable to save hash", err)
	}

	// * Save log
	size := r.pogreb.Log.Count()
	sizeBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(sizeBytes, size)
	if err := r.pogreb.Log.Put(sizeBytes, sum); err != nil {
		return nil, nil, nil, gut.Err(false, "unable to save log", err)
	}

	// * Encode base64 hash
	encoded := base64.StdEncoding.EncodeToString(sum)
	encoded = encoded[:len(encoded)-1]
	r.signature.ReplaceClean(&encoded)

	return &relativeFilePath, sum, &encoded, nil
}
