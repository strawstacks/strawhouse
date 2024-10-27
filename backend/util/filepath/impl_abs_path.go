package filepath

import "path/filepath"

func (r *Filepath) AbsPath(relativePath string) string {
	return filepath.Clean(filepath.Join(*r.config.DataRoot, relativePath))
}

func (r *Filepath) RelPath(absolutePath string) string {
	rel, err := filepath.Rel(*r.config.DataRoot, absolutePath)
	if err != nil {
		return ""
	}
	return "/" + rel
}

func (r *Filepath) CombinePath(basePath string, relativePath string) string {
	return filepath.Clean(filepath.Join(basePath, relativePath))
}
