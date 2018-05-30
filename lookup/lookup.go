package lookup

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/subchen/go-stack/fs"
)

var errNotFound = errors.New("file not found")

// default search dirs
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

// Find finds the first matched filename in search dirs
func Find(filename ...string) (string, error) {
	// lookup in dirs
	for _, dir := range dirs {
		file, err := FindAt(dir, filename...)
		if err == nil {
			return file, nil
		}
	}

	return "", errNotFound
}

// FindAt finds the first matched filename in given dir
func FindAt(dir string, filename ...string) (string, error) {
	for _, file := range filename {
		matches, err := filepath.Glob(path.Join(dir, file))
		if err == nil && len(matches) > 0 {
			return matches[0], nil
		}
	}

	return "", errNotFound
}
