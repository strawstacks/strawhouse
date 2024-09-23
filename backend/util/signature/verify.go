package signature

import (
	"encoding/base64"
	uu "github.com/bsthun/goutils"
	"strings"
	"time"
)

func (r *Signature) Verify(path string, token string) *uu.ErrorInstance {
	// * Reconstruct data
	token = strings.ReplaceAll(token, "*", "+")
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil || len(data) != 18 {
		return uu.Err(false, "Malformed token")
	}

	// Extract data
	version := data[0]
	mode := (data[1] & 0b10000000) >> 7
	depth := uint32(data[1] & 0b00111111)
	offset := (uint64(data[2]) << 32) | (uint64(data[3]) << 24) | (uint64(data[4]) << 16) | (uint64(data[5]) << 8) | uint64(data[6])
	expired := time.Unix(int64(offset), 0)

	// * Check version
	if version != 1 {
		return uu.Err(false, "Token version not supported")
	}

	// * Check expired
	if time.Now().After(expired) {
		return uu.Err(false, "Token expired")
	}

	// * Reconstruct path
	var pathValue []byte
	if mode == 0 {
		pathValue = []byte(path)
	} else {
		pathValue = extractPathSlice(path, depth)
	}

	// * Sign data
	r.Hash.Reset()
	r.Hash.Write(data[:7])
	r.Hash.Write(pathValue)
	signature := r.Hash.Sum(nil)
	copy(data[7:], signature[:11])

	// * Convert data to base64
	encodedData := base64.StdEncoding.EncodeToString(data[:])

	// * Compare token
	if token != encodedData {
		return uu.Err(false, "Invalid token")
	}

	return nil
}