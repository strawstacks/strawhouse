package name

import "regexp"

type Name struct {
	HarmfulRegexp *regexp.Regexp
	ConsecRegexp  *regexp.Regexp
}

func Init() *Name {
	harmfulRegexp := regexp.MustCompile(`[<>:"/\\|?*]`)
	consecRegexp := regexp.MustCompile(`_+`)
	return &Name{
		HarmfulRegexp: harmfulRegexp,
		ConsecRegexp:  consecRegexp,
	}
}
