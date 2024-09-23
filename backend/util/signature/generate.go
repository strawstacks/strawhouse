package signature

import (
	"encoding/base64"
	"github.com/davecgh/go-spew/spew"
	"reflect"
	"time"
	"unsafe"
)

func (r *Signature) Generate(version uint8, mode uint8, depth uint32, expired time.Time, path string) string {
	// * Construct data
	data := make([]byte, 18)

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
	r.Hash.Reset()
	r.Hash.Write(*(*[]byte)(unsafe.Pointer(&splitDataHeader)))
	r.Hash.Write(pathSlice)
	signature := r.Hash.Sum(nil)
	copy(data[7:], signature[:11])

	// * Convert data to base64
	base64buffer := make([]byte, 24)
	base64.StdEncoding.Encode(base64buffer, data)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&base64buffer))
	stringHeader := reflect.StringHeader{Data: sliceHeader.Data, Len: sliceHeader.Len}
	encoded := *(*string)(unsafe.Pointer(&stringHeader))
	ReplaceChar(&encoded, '+', '*')
	spew.Dump(encoded)
	return encoded
}
