package strawhouse

import (
	"bytes"
	"encoding/base64"
	uu "github.com/bsthun/goutils"
	"reflect"
	"time"
	"unsafe"
)

func (r *Signature) Verify(act SignatureAction, path string, attribute []byte, token string) *uu.ErrorInstance {
	// * Reconstruct data
	r.ReplaceChar(&token, '*', '+')
	data := make([]byte, 27)
	ln, err := base64.StdEncoding.Decode(data, []byte(token))
	if err != nil || ln != 27 {
		return uu.Err(false, "malformed token")
	}

	// Extract data
	version := data[0]
	mode := SignatureMode((data[1] & 0b10000000) >> 7)
	action := SignatureAction((data[1] & 0b01000000) >> 6)
	depth := uint32(data[1] & 0b00111111)
	offset := (uint64(data[2]) << 32) | (uint64(data[3]) << 24) | (uint64(data[4]) << 16) | (uint64(data[5]) << 8) | uint64(data[6])
	expired := time.Unix(int64(offset), 0)

	// * Check version
	if version != 1 {
		return uu.Err(false, "token version not supported")
	}

	// * Check action
	if act != action {
		return uu.Err(false, "invalid action")
	}

	// * Check expired
	if time.Now().After(expired) {
		return uu.Err(false, "token expired")
	}

	// * Reconstruct path
	var pathValue []byte
	if mode == SignatureModeFile {
		pathValue = []byte(path)
	} else if mode == SignatureModeDirectory {
		if action == SignatureActionUpload {
			pathValue = r.extractDirSlice(path)
		}
		if action == SignatureActionGet {
			pathValue = r.extractPathSlice(path, depth)
		}
	} else {
		uu.Fatal("invalid mode", nil)
	}

	// * Sign data
	dataHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	splitDataHeader := reflect.SliceHeader{Data: dataHeader.Data, Len: 7, Cap: 27}
	hash := r.GetHash()
	hash.Write(*(*[]byte)(unsafe.Pointer(&splitDataHeader)))
	hash.Write(pathValue)
	hash.Write(attribute)
	signature := hash.Sum(nil)
	r.PutHash(hash)

	// * Compare token
	if !bytes.Equal(data[7:], signature[:20]) {
		return uu.Err(false, "invalid token")
	}

	return nil
}
