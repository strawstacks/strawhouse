package strawhouse

import (
	"bytes"
	"encoding/base64"
	"github.com/bsthun/gut"
	"reflect"
	"time"
	"unsafe"
)

func (r *Signature) Verify(act SignatureAction, path string, token string) (string, error) {
	return r.VerifyInt(act, path, token)
}

func (r *Signature) VerifyInt(act SignatureAction, path string, token string) (string, *gut.ErrorInstance) {
	// * Reconstruct data
	if len(token) < 40 {
		return "", gut.Err(false, "token too short")
	}
	r.ReplaceUnclean(&token)

	// * Decode token
	data := make([]byte, 30)
	ln, err := base64.StdEncoding.Decode(data, []byte(token)[:40])
	if err != nil || ln != 30 {
		return "", gut.Err(false, "malformed token", err)
	}

	// Extract data
	version := data[0]
	action := SignatureAction((data[1] & 0b10000000) >> 7)
	mode := SignatureMode((data[1] & 0b01000000) >> 6)
	depth := data[1] & 0b00111110 >> 1
	recursive := data[1] & 0b00000001
	offset := (uint64(data[2]) << 32) | (uint64(data[3]) << 24) | (uint64(data[4]) << 16) | (uint64(data[5]) << 8) | uint64(data[6])
	expired := time.Unix(int64(offset), 0)

	// * Check version
	if version != 1 {
		return "", gut.Err(false, "token version not supported")
	}

	// * Check action
	if act != action {
		return "", gut.Err(false, "invalid action")
	}

	// * Check expired
	if time.Now().After(expired) {
		return "", gut.Err(false, "token expired")
	}

	// * Reconstruct path
	var pathValue []byte
	if mode == SignatureModeFile {
		pathValue = []byte(path)
	} else if mode == SignatureModeDirectory {
		pathValue = r.extractPathSlice(path, depth)
		if r.CountFixedDepth(path) < depth {
			return "", gut.Err(false, "accessing non permitted depth")
		}
		if recursive == 0 && r.CountFixedDepth(path) > depth {
			return "", gut.Err(false, "non permitted recursive")
		}
	} else {
		gut.Fatal("invalid mode", nil)
	}

	// * Reconstruct attribute
	attribute := []byte(token[40:])

	// * Sign data
	dataHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	splitDataHeader := reflect.SliceHeader{Data: dataHeader.Data, Len: 7, Cap: 7}
	hash := r.GetHash()
	hash.Write(*(*[]byte)(unsafe.Pointer(&splitDataHeader)))
	hash.Write(pathValue)
	hash.Write(attribute)
	signature := hash.Sum(nil)
	r.PutHash(hash)

	// * Compare token
	if !bytes.Equal(data[7:], signature[:23]) {
		return "", gut.Err(false, "invalid token")
	}

	return token[40:], nil
}
