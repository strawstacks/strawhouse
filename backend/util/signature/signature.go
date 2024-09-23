package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"hash"
)

type Signature struct {
	Hash hash.Hash
}

func NewSignature(key string) *Signature {
	h := hmac.New(sha256.New, []byte(key))
	return &Signature{
		Hash: h,
	}
}
