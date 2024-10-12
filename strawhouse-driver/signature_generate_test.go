package strawhouse

import (
	"bytes"
	"encoding/gob"
	"github.com/bsthun/gut"
	"testing"
	"time"
)

func Prepare() (*Signature, []byte) {
	signature := NewSignature("secret")
	attribute := &ExampleAttribute{
		UploaderId:  gut.Ptr[uint64](20),
		SessionName: gut.Ptr("abcd"),
	}
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	gob.Register(new(ExampleAttribute))
	_ = enc.Encode(attribute)
	return signature, buffer.Bytes()
}

func BenchmarkGenerate(b *testing.B) {
	signature, attr := Prepare()
	for n := 0; n < b.N; n++ {
		signature.Generate(1, SignatureModeDirectory, SignatureActionGet, 2, time.Now().Add(20*time.Minute), "/photo/2024", attr)
	}
}
