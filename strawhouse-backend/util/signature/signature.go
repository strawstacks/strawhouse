package signature

import (
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"github.com/strawstacks/strawhouse/strawhouse-driver"
)

func Init(config *config.Config) *strawhouse.Signature {
	return strawhouse.NewSignature(*config.Key)
}
