package signature

import (
	"github.com/strawstacks/strawhouse-go"
	"strawhouse-backend/common/config"
)

func Init(config *config.Config) *strawhouse.Signature {
	return strawhouse.NewSignature(*config.Key)
}
