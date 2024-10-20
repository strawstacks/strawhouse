package strawhouse

import (
	"encoding/base64"
	"github.com/bsthun/gut"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

func (r *Signature) Generate(action SignatureAction, mode SignatureMode, path string, recursive bool, expired time.Time, attribute []byte) string {
	// * Spell check path
	if path[0] != '/' {
		gut.Fatal("Path must start with /", nil)
	}
	if mode == SignatureModeDirectory && path[len(path)-1] != '/' {
		gut.Fatal("Path must end with /", nil)
	}

	// * Construct data
	data := make([]byte, 30)

	// * Add version 1 byte
	data[0] = uint8(1)

	// * Add action 1 bit
	if action == SignatureActionGet {
		data[1] &= 0b01111111
	} else if action == SignatureActionUpload {
		data[1] |= 0b10000000
	} else {
		gut.Fatal("Invalid action: "+strconv.Itoa(int(action)), nil)
	}

	// * Add metadata 1 bit
	if mode == SignatureModeFile {
		data[1] &= 0b10111111
	} else if mode == SignatureModeDirectory {
		data[1] |= 0b01000000
	} else {
		gut.Fatal("Invalid mode: "+strconv.Itoa(int(mode)), nil)
	}

	// * Add fixed depth 5 bits
	depth := r.CountFixedDepth(path)
	if depth > 0b11111 {
		depth = 0b11111
	}
	data[1] |= depth << 1

	// * Add recursive 1 bit
	if recursive {
		data[1] |= 0b00000001
	} else {
		data[1] &= 0b11111110
	}

	// * Add expired time 5 bytes
	offset := uint64(expired.Unix())
	if offset > 0xFFFFFFFFFF {
		offset = 0xFFFFFFFFFF
	}
	data[2] = byte(offset >> 32)
	data[3] = byte(offset >> 24)
	data[4] = byte(offset >> 16)
	data[5] = byte(offset >> 8)
	data[6] = byte(offset)

	// Add salt 3 byte
	gut.Rand.Read(data[7:10])

	// * Sign data
	dataHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	splitDataHeader := reflect.SliceHeader{Data: dataHeader.Data, Len: 10, Cap: 10}
	hash := r.GetHash()
	hash.Write(*(*[]byte)(unsafe.Pointer(&splitDataHeader)))
	hash.Write([]byte(path))
	hash.Write(attribute)
	signature := hash.Sum(nil)
	r.PutHash(hash)
	copy(data[10:], signature[:20])

	// * Encode data
	attributeLength := base64.StdEncoding.EncodedLen(len(attribute))
	buffer := make([]byte, 40+attributeLength)
	base64.StdEncoding.Encode(buffer, data)

	attributeBuffer := make([]byte, attributeLength)
	base64.StdEncoding.Encode(attributeBuffer, attribute)

	// * Merge data and attribute
	copy(buffer[40:], attributeBuffer[:])
	encoded := string(buffer[:])
	r.ReplaceClean(&encoded)

	return encoded
}
