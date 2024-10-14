package strawhouse

import (
	"reflect"
	"unsafe"
)

type ExampleAttribute struct {
	UploaderId  *uint64
	SessionName *string
}

func (r *Signature) extractPathSlice(path string, depth uint8) []byte {
	count := int(depth)
	for index := 0; index < len(path); index++ {
		if path[index] == '/' {
			count--
			if count == 0 {
				return unsafe.Slice(unsafe.StringData(path), index+1)
			}
		}
	}

	return unsafe.Slice(unsafe.StringData(path), len(path))
}

func (r *Signature) CountFixedDepth(path string) uint8 {
	depth := uint8(0)
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			depth++
			if depth == 0b111 {
				break
			}
		}
	}
	return depth
}

func (r *Signature) ReplaceClean(str *string) {
	byteSlice := (*[]byte)(unsafe.Pointer(&reflect.StringHeader{
		Data: (*reflect.StringHeader)(unsafe.Pointer(str)).Data,
		Len:  len(*str),
	}))
	for i := 0; i < len(*byteSlice); i++ {
		if (*byteSlice)[i] == '+' {
			(*byteSlice)[i] = '-'
		}
		if (*byteSlice)[i] == '/' {
			(*byteSlice)[i] = '_'
		}
	}
}

func (r *Signature) ReplaceUnclean(str *string) {
	byteSlice := (*[]byte)(unsafe.Pointer(&reflect.StringHeader{
		Data: (*reflect.StringHeader)(unsafe.Pointer(str)).Data,
		Len:  len(*str),
	}))
	for i := 0; i < len(*byteSlice); i++ {
		if (*byteSlice)[i] == '-' {
			(*byteSlice)[i] = '+'
		}
		if (*byteSlice)[i] == '_' {
			(*byteSlice)[i] = '/'
		}
	}
}
