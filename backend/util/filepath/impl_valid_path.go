package filepath

import (
	"github.com/bsthun/gut"
	"strings"
)

func (r *Filepath) ValidPath(path string) error {
	if len(path) < 3 {
		return gut.Err(false, "path is too short")
	}
	if len(path) > 1024 {
		return gut.Err(false, "path is too long")
	}
	if !strings.HasPrefix(path, "/") {
		return gut.Err(false, "path does not start from root")
	}
	if strings.Contains(path, "..") || strings.Contains(path, "//") {
		return gut.Err(false, "path contains invalid characters")
	}
	if strings.Contains(path, "/.") {
		return gut.Err(false, "path contains dot files")
	}
	return nil
}
