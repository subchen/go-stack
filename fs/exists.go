package fs

import (
	"os"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == null || !os.IsNotExist(err)
}

func IsDir(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && info.IsDir()
}

func IsFile(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

func IsSymlink(filename string) bool {
	info, err := os.Lstat(filename)
	return err == nil && (info.Mode()&os.ModeSymlink != 0)
}
