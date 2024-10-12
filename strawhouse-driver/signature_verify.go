package strawhouse

import (
	"bytes"
	"encoding/base64"
	"github.com/bsthun/gut"
	"reflect"
	"time"
	"unsafe"
)

func (r *Signature) Verify(act SignatureAction, path string, attribute []byte, token string) error {
	return r.VerifyInt(act, path, attribute, token)
}

func (r *Signature) VerifyInt(act SignatureAction, path string, attribute []byte, token string) *gut.ErrorInstance {
	// * Reconstruct data
	r.ReplaceUnclean(&token)
	data := make([]byte, 27)
	ln, err := base64.StdEncoding.Decode(data, []byte(token))
	if err != nil || ln != 27 {
		return gut.Err(false, "malformed token")
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
		return gut.Err(false, "token version not supported")
	}

	// * Check action
	if act != action {
		return gut.Err(false, "invalid action")
	}

	// * Check expired
	if time.Now().After(expired) {
		return gut.Err(false, "token expired")
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
		gut.Fatal("invalid mode", nil)
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
		return gut.Err(false, "invalid token")
	}

	return nil
}
