package strawhouse

import (
	"github.com/bsthun/gut"
	"testing"
	"time"
)

func Prepare() (*Signature, string) {
	signature := NewSignature("secret")
	attribute := gut.Random(gut.RandomSet.MixedAlphaNum, 16)
	return signature, *attribute
}

func BenchmarkGenerate(b *testing.B) {
	signature, attr := Prepare()
	for n := 0; n < b.N; n++ {
		signature.Generate(SignatureActionGet, SignatureModeDirectory, "/photo/2024/", true, time.Now().Add(20*time.Minute), []byte(attr))
	}
}
