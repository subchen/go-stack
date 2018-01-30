package fs

import (
	"os"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

func DirExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && info.IsDir()
}
