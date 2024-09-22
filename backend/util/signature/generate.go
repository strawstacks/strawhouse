package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"time"
)

func Generate(version uint8, mode uint8, depth uint32, expired time.Time, path string, key string) string {
	// * Construct data
	var data [18]byte

	// * Add version 1 byte
	data[0] = version

	// * Add mode 1 byte
	var pathSlice []byte
	if mode == 0 {
		data[1] &= 0b01111111
		pathSlice = []byte(path)
	} else {
		data[1] |= 0b10000000
		pathSlice = extractPathSlice(path, depth)

		// * Add depth 6 bits
		if depth > 63 {
			depth = 63
		}
		data[1] |= byte(depth)
	}

	// * Add expired time 5 bytes
	if expired.Before(startTime) {
		expired = startTime
	}
	offset := uint64(expired.Unix() - startTime.Unix())
	if offset > 0xFFFFFFFFFF {
		offset = 0xFFFFFFFFFF
	}

	// * Add 5 bytes offset
	data[2] = byte(offset >> 32)
	data[3] = byte(offset >> 24)
	data[4] = byte(offset >> 16)
	data[5] = byte(offset >> 8)
	data[6] = byte(offset)

	// * Sign data
	hmacHash := hmac.New(sha256.New, []byte(key))
	hmacHash.Write(data[:7])
	hmacHash.Write(pathSlice)
	signature := hmacHash.Sum(nil)
	copy(data[7:], signature[:11])

	// * Convert data to base64
	encodedData := base64.StdEncoding.EncodeToString(data[:])
	encodedData = strings.ReplaceAll(encodedData, "+", "*")

	return encodedData
}
