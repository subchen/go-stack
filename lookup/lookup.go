package lookup

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/subchen/go-stack/fs"
)

var dirs = []string{
	fs.GetProcessCWD(),
	fs.GetProcessBinDir(),
}

// AddSearchDir add dir to lookup
func AddSearchDir(dirname string) {
	dir, err := filepath.Abs(dirname)
	if err != nil {
		panic(err)
	}
	dirs = append(dirs, dir)
}

// AddGopkgDir add package dir under GOPATH to lookup
func AddGopkgDir(pkgname string) {
	gopath := os.Getenv("GOPATH")
	if gopath != "" {
		for _, root := range strings.Split(gopath, ":") {
			dir := filepath.Join(root, pkgname)
			dirs = append(dirs, dir)
		}
	}
}

// GetFile returns full filename in relative path
func GetFile(filename string) string {
	// skip if abs
	if filepath.IsAbs(filename) {
		if !fs.IsFile(filename) {
			panic(fmt.Sprintf("file not found: %v", filename))
		}
		return filename
	}

	// lookup in dirs
	for _, dir := range dirs {
		if f := filepath.Join(dir, filename); fs.IsFile(f) {
			return f
		}
	}

	panic(fmt.Sprintf("file not found: %v", filename))
}
