package gstack

import (
	"path/filepath"
	"strings"
)

func PathGetNameWithoutExt(file string) string {
	name := filepath.Base(file)
	ext := filepath.Ext(name)
	return strings.TrimSuffix(name, ext)
}
