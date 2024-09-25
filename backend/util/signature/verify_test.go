package signature

import (
	uu "github.com/bsthun/goutils"
	"github.com/go-playground/assert/v2"
	"github.com/strawstacks/strawhouse/backend/type/enum"
	"testing"
	"time"
)

func TestVerify(t *testing.T) {
	signature, attr := Prepare()

	token := signature.Generate(1, enum.SignatureModeDirectory, enum.SignatureActionUpload, 2, time.Now().Add(20*time.Minute), "/photo/2024/0628-sao-kmutt-seminar", attr)
	assert.Equal(t, signature.Verify(enum.SignatureActionUpload, "/photo/2023/1127-sit-dday", attr, token), uu.Err(false, "invalid token"))
	assert.Equal(t, signature.Verify(enum.SignatureActionUpload, "/photo/2024/0528-dsi-iot-workshop", attr, token), nil)
	assert.Equal(t, signature.Verify(enum.SignatureActionGet, "/photo/2024/0528-dsi-iot-workshop", attr, token), uu.Err(false, "invalid action"))
	assert.Equal(t, signature.Verify(enum.SignatureActionUpload, "/photo/2024/0528-dsi-iot-workshop", []byte{0x21}, token), uu.Err(false, "invalid token"))
	assert.Equal(t, signature.Verify(enum.SignatureActionUpload, "/photo/2024/hello.jpg", attr, token), nil)
	assert.Equal(t, signature.Verify(enum.SignatureActionGet, "/photo/2024/hello2.jpg", attr, token), uu.Err(false, "invalid action"))

	token = signature.Generate(1, enum.SignatureModeDirectory, enum.SignatureActionUpload, 3, time.Now().Add(-2*time.Second), "/photo/2023/1217-ctouting1", attr)
	assert.Equal(t, signature.Verify(enum.SignatureActionUpload, "/photo/2023/1217-ctouting1/a2.png", attr, token), uu.Err(false, "token expired"))

	token = signature.Generate(1, enum.SignatureModeFile, enum.SignatureActionGet, 2, time.Now().Add(2*time.Second), "/photo/2023/1217-ctouting1/1.jpg", nil)
	assert.Equal(t, signature.Verify(enum.SignatureActionGet, "/photo/2023/1217-ctouting1/1.jpg", nil, token), nil)
	assert.Equal(t, signature.Verify(enum.SignatureActionGet, "/photo/2023/1217-ctouting1/1.jpg", attr, token), uu.Err(false, "invalid token"))
	assert.Equal(t, signature.Verify(enum.SignatureActionGet, "/photo/2023/1217-ctouting1/2.jpg", attr, token), uu.Err(false, "invalid token"))
	assert.Equal(t, signature.Verify(enum.SignatureActionUpload, "/photo/2023/1217-ctouting1/1.jpg", attr, token), uu.Err(false, "invalid action"))

	token = signature.Generate(1, enum.SignatureModeFile, enum.SignatureActionGet, 2, time.Now().Add(2*time.Second), "/photo/2023/1217-ctouting1/1.jpg", []byte{0x21})
	assert.Equal(t, signature.Verify(enum.SignatureActionGet, "/photo/2023/1217-ctouting1/1.jpg", []byte{0x21}, token), nil)
	assert.Equal(t, signature.Verify(enum.SignatureActionGet, "/photo/2023/1217-ctouting1/1.jpg", nil, token), uu.Err(false, "invalid token"))
	assert.Equal(t, signature.Verify(enum.SignatureActionGet, "/photo/2023/1217-ctouting1/1.jpg", attr, token), uu.Err(false, "invalid token"))
	assert.Equal(t, signature.Verify(enum.SignatureActionUpload, "/photo/2023/1217-ctouting1/1.jpg", attr, token), uu.Err(false, "invalid action"))
}

func BenchmarkVerify(b *testing.B) {
	signature, attr := Prepare()
	token := signature.Generate(1, enum.SignatureModeDirectory, enum.SignatureActionUpload, 2, time.Now().Add(20*time.Minute), "/photo/2024/", attr)
	for n := 0; n < b.N; n++ {
		_ = signature.Verify(enum.SignatureActionUpload, "/photo/2024/1.jpg", attr, token)
	}
}
