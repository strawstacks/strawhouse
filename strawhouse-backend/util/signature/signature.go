package signature

import (
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"github.com/strawstacks/strawhouse/strawhouse-driver"
	"reflect"
	"unsafe"
)

func Init(config *config.Config) *strawhouse.Signature {
	return strawhouse.NewSignature(*config.Key)
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
