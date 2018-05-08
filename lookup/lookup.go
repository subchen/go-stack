package lookup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var dirs = []string{
	getWorkingdir(),
	getProcessdir(),
}

// AddDir add dir to lookup
func AddDir(dirname string) {
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
		if !exists(filename) {
			panic(fmt.Sprintf("file not found: %v", filename))
		}
		return filename
	}

	// lookup in dirs
	for _, dir := range dirs {
		if f := filepath.Join(dir, filename); exists(f) {
			return f
		}
	}

	panic(fmt.Sprintf("file not found: %v", filename))
}

// working dir
func getWorkingdir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}

// process bin dir
func getProcessdir() string {
	app := os.Args[0]
	if filepath.Base(app) == app {
		file, err := exec.LookPath(app)
		if err != nil {
			panic(err)
		}
		app = file
	}
	dir, err := filepath.Abs(filepath.Dir(app))
	if err != nil {
		panic(err)
	}
	return dir
}

// file exists
func exists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
