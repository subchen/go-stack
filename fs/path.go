package fs

import (
	"os"
	"path/filepath"
)

// GetProcessBinDir returns process binary directory
func GetProcessBinDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
}

// GetProcessPWD returns process working directory
func GetProcessPWD() string {
	dir := os.Getenv("PWD")
	if dir == "" {
		dir = "."
	}
	return dir
}
