package strawhouse

import (
	"crypto/hmac"
	"crypto/sha256"
	"hash"
	"sync"
)

type Signature struct {
	HashPool *sync.Pool
}

func NewSignature(key string) *Signature {
	hashPool := &sync.Pool{
		New: func() any {
			return hmac.New(sha256.New, []byte(key))
		},
	}
	return &Signature{
		HashPool: hashPool,
	}
}

func (r *Signature) GetHash() hash.Hash {
	return r.HashPool.Get().(hash.Hash)
}

func (r *Signature) PutHash(hash hash.Hash) {
	hash.Reset()
	r.HashPool.Put(hash)
}