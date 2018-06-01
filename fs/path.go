package fs

import (
	"path/filepath"
	"strings"
)

// BasenameWithoutExt returns file basename without ext
func BasenameWithoutExt(file string) string {
	name := filepath.Base(file)
	ext := filepath.Ext(name)
	return strings.TrimSuffix(name, ext)
}
