package gstack

import (
	"os"
	"path/filepath"
)

// ProcessGetBinDir returns process binary directory
func ProcessGetBinDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
}

// ProcessGetPWD returns process binary directory
func ProcessGetPWD() string {
	dir := os.Getenv("PWD")
	if dir == "" {
		dir = "."
	}
	return dir
}
