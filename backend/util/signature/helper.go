package signature

import (
	"reflect"
	"unsafe"
)

func extractPathSlice(path string, depth uint32) []byte {
	if depth == 0 {
		return unsafe.Slice(unsafe.StringData(path), 1)
	}

	count := int(depth)
	for index := 1; index < len(path); index++ {
		if path[index] == '/' {
			count--
			if count <= 0 {
				return unsafe.Slice(unsafe.StringData(path), index)
			}
		}
	}

	return unsafe.Slice(unsafe.StringData(path), 1)
}

func ReplaceChar(str *string, oldChar, newChar rune) {
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
