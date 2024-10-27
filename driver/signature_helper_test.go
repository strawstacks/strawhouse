package strawhouse

import (
	"testing"
)

func BenchmarkExtractPathSlice(b *testing.B) {
	signature, _ := Prepare()

	// run function b.N times
	for n := 0; n < b.N; n++ {
		signature.extractPathSlice("/api/v1/health", 2)
	}
}
