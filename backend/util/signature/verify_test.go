package signature

import (
	"backend/type/enum"
	uu "github.com/bsthun/goutils"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

func TestVerify(t *testing.T) {
	signature, attr := Prepare()

	token := signature.Generate(1, enum.SignatureModeDirectory, enum.SignatureActionGet, 2, time.Now().Add(20*time.Minute), "/photo/2024/0628-sao-kmutt-seminar", attr)
	assert.NotEqual(t, signature.Verify(enum.SignatureActionGet, "/photo/2023/1127-sit-dday", attr, token), nil)
	assert.Equal(t, signature.Verify(enum.SignatureActionGet, "/photo/2024/0528-dsi-iot-workshop", attr, token), nil)
	assert.NotEqual(t, signature.Verify(enum.SignatureActionUpload, "/photo/2024/0528-dsi-iot-workshop", attr, token), nil)
	assert.NotEqual(t, signature.Verify(enum.SignatureActionGet, "/photo/2024/0528-dsi-iot-workshop", []byte{0x21}, token), nil)
	assert.Equal(t, signature.Verify(enum.SignatureActionGet, "/photo/2024/hello.jpg", attr, token), nil)

	token = signature.Generate(1, enum.SignatureModeDirectory, enum.SignatureActionUpload, 3, time.Now().Add(-2*time.Second), "/photo/2023/1217-ctouting1", attr)
	assert.Equal(t, signature.Verify(enum.SignatureActionUpload, "/photo/2023/1217-ctouting1", attr, token), uu.Err(false, "Token expired"))
}

func BenchmarkVerify(b *testing.B) {
	signature, attr := Prepare()
	token := signature.Generate(1, enum.SignatureModeDirectory, enum.SignatureActionGet, 2, time.Now().Add(20*time.Minute), "/photo/2024/", attr)
	for n := 0; n < b.N; n++ {
		_ = signature.Verify(enum.SignatureActionGet, "/photo/2024/1.jpg", attr, token)
	}
}
