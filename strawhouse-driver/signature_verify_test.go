package strawhouse

import (
	"github.com/bsthun/gut"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

func TestVerify(t *testing.T) {
	signature, attr := Prepare()

	token := signature.Generate(SignatureActionGet, SignatureModeFile, "/photo/2023/1217-ctouting1/1.jpg", false, time.Now().Add(2*time.Second), []byte(attr))

	attribute, err := signature.VerifyInt(SignatureActionGet, "/photo/2023/1217-ctouting1/1.jpg", token)
	assert.Equal(t, err, nil)
	assert.Equal(t, attribute, []byte(attr))

	attribute, err = signature.VerifyInt(SignatureActionGet, "/photo/2023/1217-ctouting1/1.jpg2", token)
	assert.Equal(t, err, gut.Err(false, "invalid token"))

	attribute, err = signature.VerifyInt(SignatureActionUpload, "/photo/2023/1217-ctouting1/1.jpg", token)
	assert.Equal(t, err, gut.Err(false, "invalid action"))

	token = signature.Generate(SignatureActionUpload, SignatureModeDirectory, "/photo/2024/0628-sao-kmutt-seminar/", false, time.Now().Add(20*time.Minute), []byte(attr))

	attribute, err = signature.VerifyInt(SignatureActionUpload, "/photo/2024/0628-sao-kmutt-seminar/", token)
	assert.Equal(t, err, nil)
	assert.Equal(t, attribute, []byte(attr))

	attribute, err = signature.VerifyInt(SignatureActionUpload, "/photo/2024/0628-sao-kmutt-seminar/2.png", token)
	assert.Equal(t, err, nil)

	attribute, err = signature.VerifyInt(SignatureActionUpload, "/photo/2023/1127-sit-dday/", token)
	assert.Equal(t, err, gut.Err(false, "invalid token"))

	attribute, err = signature.VerifyInt(SignatureActionUpload, "/photo/2024/0628-sao-kmutt-seminar", token)
	assert.Equal(t, err, gut.Err(false, "accessing non permitted depth"))

	attribute, err = signature.VerifyInt(SignatureActionUpload, "/photo/2024/hello.jpg", token)
	assert.Equal(t, err, gut.Err(false, "accessing non permitted depth"))

	attribute, err = signature.VerifyInt(SignatureActionUpload, "/photo/2024/0628-sao-kmutt-seminar/img/2.png", token)
	assert.Equal(t, err, gut.Err(false, "non permitted recursive"))

	attribute, err = signature.VerifyInt(SignatureActionGet, "/photo/2024/hello2.jpg", token)
	assert.Equal(t, err, gut.Err(false, "invalid action"))

	token = signature.Generate(SignatureActionGet, SignatureModeDirectory, "/photo/2024/1127-sit-dday/", true, time.Now().Add(20*time.Minute), []byte(attr))

	attribute, err = signature.VerifyInt(SignatureActionGet, "/photo/2024/1127-sit-dday/_DSC5090.jpg", token)
	assert.Equal(t, err, nil)

	attribute, err = signature.VerifyInt(SignatureActionGet, "/photo/2024/1127-sit-dday/img/_DSC5090.jpg", token)
	assert.Equal(t, err, nil)

	attribute, err = signature.VerifyInt(SignatureActionGet, "/photo/2024/1127-sit-dday", token)
	assert.Equal(t, err, gut.Err(false, "accessing non permitted depth"))

	token = signature.Generate(SignatureActionUpload, SignatureModeDirectory, "/photo/2023/1217-ctouting1/", false, time.Now().Add(-2*time.Second), []byte(attr))

	attribute, err = signature.VerifyInt(SignatureActionUpload, "/photo/2023/1217-ctouting1/a2.png", token)
	assert.Equal(t, err, gut.Err(false, "token expired"))
}

func BenchmarkVerify(b *testing.B) {
	signature, attr := Prepare()
	token := signature.Generate(SignatureActionUpload, SignatureModeDirectory, "/photo/2024/", true, time.Now().Add(20*time.Minute), []byte(attr))
	for n := 0; n < b.N; n++ {
		_, _ = signature.VerifyInt(SignatureActionUpload, "/photo/2024/1.jpg", token)
	}
}
