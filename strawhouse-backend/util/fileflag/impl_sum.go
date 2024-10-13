package fileflag

import (
	"bytes"
	"github.com/bsthun/gut"
	"github.com/pkg/xattr"
)

func (r *Fileflag) SumSet(relativePath string, sum []byte) *gut.ErrorInstance {
	// * Convert path
	absolutePath := r.filepath.AbsPath(relativePath)

	// * Sign checksum
	hash := r.signature.GetHash()
	hash.Write(sum)
	signedSum := hash.Sum(nil)
	r.signature.PutHash(hash)

	// * Set file attributes
	err := xattr.Set(absolutePath, xattrSumTag, []byte(string(sum)+":"+string(signedSum)))
	if err != nil {
		return gut.Err(false, "unable to set file sum attributes", err)
	}

	return nil
}

func (r *Fileflag) SumGet(relativePath string) ([]byte, *gut.ErrorInstance) {
	// * Convert path
	absolutePath := r.filepath.AbsPath(relativePath)

	// * Get file attributes
	attr, err := xattr.Get(absolutePath, xattrSumTag)
	if err != nil {
		return nil, gut.Err(false, "unable to get file sum attributes", err)
	}

	// * Split sum and signed sum
	parts := bytes.SplitN(attr, []byte(":"), 2)
	if len(parts) != 2 {
		return nil, gut.Err(false, "invalid sum attributes format", nil)
	}
	sum := parts[0]
	signedSum := parts[1]

	// * Verify signed checksum
	hash := r.signature.GetHash()
	hash.Write(sum)
	expectedSignedSum := hash.Sum(nil)
	r.signature.PutHash(hash)

	// * Compare signed checksum
	if !bytes.Equal(signedSum, expectedSignedSum) {
		return nil, gut.Err(false, "invalid signed checksum", nil)
	}

	return sum, nil
}
