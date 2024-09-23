package signature

import (
	"backend/common/config"
	"crypto/hmac"
	"crypto/sha256"
	"hash"
)

type Signature struct {
	Hash hash.Hash
}

func Init(config *config.Config) *Signature {
	return New(*config.Clients[0].Key)
}

func New(key string) *Signature {
	h := hmac.New(sha256.New, []byte(key))
	return &Signature{
		Hash: h,
	}
}
