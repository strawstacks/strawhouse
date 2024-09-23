package signature

import (
	"testing"
	"time"
)

func BenchmarkGenerate(b *testing.B) {
	signature := NewSignature("secret")

	for n := 0; n < b.N; n++ {
		signature.Generate(1, 2, 1, time.Now().Add(20*time.Minute), "/photo/2024")
	}
}
