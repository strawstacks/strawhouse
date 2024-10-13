package filepath

import (
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"regexp"
)

type Filepath struct {
	config *config.Config
	val    *FilapathVal
}

type FilapathVal struct {
	harmfulRegexp *regexp.Regexp
	consecRegexp  *regexp.Regexp
}

func Init(config *config.Config) *Filepath {
	harmfulRegexp := regexp.MustCompile(`[<>:"/\\|?*]`)
	consecRegexp := regexp.MustCompile(`_+`)
	return &Filepath{
		config: config,
		val: &FilapathVal{
			harmfulRegexp: harmfulRegexp,
			consecRegexp:  consecRegexp,
		},
	}
}
