package signature

import (
	"backend/type/enum"
	"bytes"
	"encoding/base64"
	uu "github.com/bsthun/goutils"
	"reflect"
	"time"
	"unsafe"
)

func (r *Signature) Verify(act enum.SignatureAction, path string, attribute []byte, token string) *uu.ErrorInstance {
	// * Reconstruct data
	ReplaceChar(&token, '*', '+')
	data := make([]byte, 27)
	ln, err := base64.StdEncoding.Decode(data, []byte(token))
	if err != nil || ln != 27 {
		return uu.Err(false, "Malformed token")
	}

	// Extract data
	version := data[0]
	mode := enum.SignatureMode((data[1] & 0b10000000) >> 7)
	action := enum.SignatureAction((data[1] & 0b01000000) >> 6)
	depth := uint32(data[1] & 0b00111111)
	offset := (uint64(data[2]) << 32) | (uint64(data[3]) << 24) | (uint64(data[4]) << 16) | (uint64(data[5]) << 8) | uint64(data[6])
	expired := time.Unix(int64(offset), 0)

	// * Check version
	if version != 1 {
		return uu.Err(false, "Token version not supported")
	}

	// * Check action
	if act != action {
		return uu.Err(false, "Invalid action")
	}

	// * Check expired
	if time.Now().After(expired) {
		return uu.Err(false, "Token expired")
	}

	// * Reconstruct path
	var pathValue []byte
	if mode == enum.SignatureModeFile {
		pathValue = []byte(path)
	} else if mode == enum.SignatureModeDirectory {
		pathValue = extractPathSlice(path, depth)
	} else {
		uu.Fatal("Invalid mode", nil)
	}

	// * Sign data
	dataHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	splitDataHeader := reflect.SliceHeader{Data: dataHeader.Data, Len: 7, Cap: 27}
	r.Hash.Reset()
	r.Hash.Write(*(*[]byte)(unsafe.Pointer(&splitDataHeader)))
	r.Hash.Write(pathValue)
	if action == enum.SignatureActionUpload {
		r.Hash.Write(attribute)
	}
	signature := r.Hash.Sum(nil)

	// * Compare token
	if !bytes.Equal(data[7:], signature[:20]) {
		return uu.Err(false, "Invalid token")
	}

	return nil
}
