package signature

import (
	"testing"
	"time"
)

func BenchmarkVerify(b *testing.B) {
	signature := New("secret")
	token := signature.Generate(1, 2, 1, time.Now().Add(20*time.Minute), "/photo/2024")
	for n := 0; n < b.N; n++ {
		_ = signature.Verify("/photo/1.jpg", token)
	}
}
