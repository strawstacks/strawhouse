package filepath

import (
	"regexp"
	"strawhouse-backend/common/config"
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
