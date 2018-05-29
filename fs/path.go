package fs

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetProcessBinDir returns process binary directory
func GetProcessBinDir() string {
	app := os.Args[0]
	if filepath.Base(app) == app {
		app, _ = exec.LookPath(app)
	}
	dir, err := filepath.Abs(filepath.Dir(app))
	if err != nil {
		panic(err)
	}
	return dir
}

// GetProcessCWD returns process working directory
func GetProcessCWD() string {
	dir, _ := os.Getwd()
	if dir == "" {
		dir = "."
	}
	return dir
}

// BasenameWithoutExt returns file basename without ext
func BasenameWithoutExt(file string) string {
	name := filepath.Base(file)
	ext := filepath.Ext(name)
	return strings.TrimSuffix(name, ext)
}
