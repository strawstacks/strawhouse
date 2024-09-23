package signature

import (
	"testing"
)

func BenchmarkExtractPathSlice(b *testing.B) {
	// run function b.N times
	for n := 0; n < b.N; n++ {
		extractPathSlice("/api/v1/health", 2)
	}
}
