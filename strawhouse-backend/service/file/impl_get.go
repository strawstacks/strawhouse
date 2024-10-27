package file

import (
	"bytes"
	"crypto/sha256"
	"github.com/bsthun/gut"
	"io"
	"os"
)

func (r *Service) Get(path string, writer io.Writer) *gut.ErrorInstance {
	// * Construct absolute path
	absolutePath := r.filepath.AbsPath(path)

	// * Check file corruption
	if er := r.fileflag.Corrupted(path); er != nil {
		return er
	}

	// * Check signed checksum
	sum, er := r.fileflag.SumGet(path)
	if er != nil {
		return er
	}

	// * Check file path
	val, err := r.pogreb.Sum.Get(sum)
	if err != nil || val == nil {
		return gut.Err(false, "file record not found")
	}
	if !bytes.Equal(val, []byte(path)) {
		return gut.Err(false, "file path mismatch")
	}

	// * Open the file
	file, err := os.Open(absolutePath)
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
			_, err := writer.Write(buffer[:n])
			if err != nil {
				return gut.Err(false, err.Error())
			}
		}
	}

	// * Compare content hash and xattr hash
	if !bytes.Equal(hash.Sum(nil), sum) {
		if er := r.fileflag.CorruptedSet(path, true); er != nil {
			return er
		}
		return gut.Err(false, "invalid file hash")
	}

	return nil
}
