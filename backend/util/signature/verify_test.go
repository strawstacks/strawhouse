package signature

import (
	"backend/type/enum"
	"testing"
	"time"
)

func BenchmarkVerify(b *testing.B) {
	signature, attr := Prepare()
	token := signature.Generate(1, enum.SignatureModeDirectory, enum.SignatureActionGet, 2, time.Now().Add(20*time.Minute), "/photo/2024", attr)
	for n := 0; n < b.N; n++ {
		_ = signature.Verify(enum.SignatureActionGet, "/photo/1.jpg", attr, token)
	}
}
