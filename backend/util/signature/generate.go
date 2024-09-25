package signature

import (
	"encoding/base64"
	uu "github.com/bsthun/goutils"
	"github.com/strawstacks/strawhouse/backend/type/enum"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

func (r *Signature) Generate(version uint8, mode enum.SignatureMode, action enum.SignatureAction, depth uint32, expired time.Time, path string, attribute []byte) string {
	// * Construct data
	data := make([]byte, 27)

	// * Add version 1 byte
	data[0] = version

	// * Add metadata 1 byte
	var pathSlice []byte
	if mode == enum.SignatureModeFile {
		data[1] &= 0b01111111
		pathSlice = []byte(path)
	} else if mode == enum.SignatureModeDirectory {
		data[1] |= 0b10000000
		pathSlice = extractPathSlice(path, depth)

		// * Add depth 6 bits
		if depth > 63 {
			depth = 63
		}
		data[1] |= byte(depth)
	} else {
		uu.Fatal("Invalid mode: "+strconv.Itoa(int(mode)), nil)
	}
	if action == enum.SignatureActionGet {
		data[1] &= 0b10111111
	} else if action == enum.SignatureActionUpload {
		data[1] |= 0b01000000
	} else {
		uu.Fatal("Invalid action: "+strconv.Itoa(int(action)), nil)
	}

	// * Add expired time 5 bytes
	offset := uint64(expired.Unix())
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
	dataHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	splitDataHeader := reflect.SliceHeader{Data: dataHeader.Data, Len: 7, Cap: 18}
	hash := r.GetHash()
	hash.Write(*(*[]byte)(unsafe.Pointer(&splitDataHeader)))
	hash.Write(pathSlice)
	hash.Write(attribute)
	signature := hash.Sum(nil)
	r.PutHash(hash)
	copy(data[7:], signature[:20])

	// * Convert data to base64
	base64buffer := make([]byte, 36)
	base64.StdEncoding.Encode(base64buffer, data)
	encoded := string(base64buffer[:])
	ReplaceChar(&encoded, '+', '*')
	return encoded
}
