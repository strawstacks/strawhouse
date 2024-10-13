package filepath

import (
	"strings"
	"unicode"
)

func (r *Filepath) BaseName(filepath string) string {
	filepath = r.val.harmfulRegexp.ReplaceAllString(filepath, "_")
	filepath = strings.Trim(filepath, ".")
	filepath = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return '_'
		}
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, filepath)
	filepath = r.val.consecRegexp.ReplaceAllString(filepath, "_")
	return filepath
}
