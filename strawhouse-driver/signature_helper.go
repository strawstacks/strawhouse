package strawhouse

import (
	"reflect"
	"unsafe"
)

type ExampleAttribute struct {
	UploaderId  *uint64
	SessionName *string
}

func (r *Signature) extractPathSlice(path string, depth uint32) []byte {
	count := int(depth)
	for index := 0; index < len(path); index++ {
		if path[index] == '/' {
			count--
			if count < 0 {
				return unsafe.Slice(unsafe.StringData(path), index+1)
			}
		}
	}

	return unsafe.Slice(unsafe.StringData(path), len(path))
}

func (r *Signature) extractDirSlice(path string) []byte {
	for index := len(path) - 1; index >= 0; index-- {
		if path[index] == '/' {
			return unsafe.Slice(unsafe.StringData(path), index+1)
		}
	}

	return nil
}

func (r *Signature) ReplaceChar(str *string, oldChar, newChar rune) {
	byteSlice := (*[]byte)(unsafe.Pointer(&reflect.StringHeader{
		Data: (*reflect.StringHeader)(unsafe.Pointer(str)).Data,
		Len:  len(*str),
	}))

	// Iterate through the byte slice and replace the old character with the new one
	for i := 0; i < len(*byteSlice); i++ {
		if (*byteSlice)[i] == byte(oldChar) {
			(*byteSlice)[i] = byte(newChar)
		}
	}
}
