package process

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Name returns process name
func Name() string {
	app := os.Args[0]
	name := filepath.Base(app)
	ext := filepath.Ext(name)
	return strings.TrimSuffix(name, ext)
}

// Dir returns process binary directory
func Dir() string {
	app := os.Args[0]
	if !strings.ContainsAny(app, `/\`) {
		app, _ = exec.LookPath(app)
	}
	dir, _ := filepath.Abs(filepath.Dir(app))
	return dir
}

// GetCwd returns process current working directory
func GetCwd() string {
	dir, _ := os.Getwd()
	if dir == "" {
		dir = "."
	}
	return dir
}
