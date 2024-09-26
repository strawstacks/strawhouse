package name

import (
	"strings"
	"unicode"
)

func (r *Name) BaseName(filepath string) string {
	filepath = r.HarmfulRegexp.ReplaceAllString(filepath, "_")
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
	filepath = r.ConsecRegexp.ReplaceAllString(filepath, "_")
	return filepath
}
