package findup

import (
	"errors"
	"os"
	"path"
	"path/filepath"
)

// Find finds the first filename matching in the current directory or the nearest ancestor directory up to root.
func Find(filename ...string) (string, error) {
	wd, _ := os.Getwd()
	return lookup(wd, filename)
}

// FindInDir finds the first filename matching in the directory or the nearest ancestor directory up to root.
func FindInDir(dir string, filename ...string) (string, error) {
	wd, _ := filepath.Abs(dir)
	return lookup(wd, filename)
}

func lookup(basepath string, filenames []string) (string, error) {
	for _, file := range filenames {
		matches, err := filepath.Glob(path.Join(basepath, file))
		if len(matches) != 0 {
			return matches[0], err
		}
	}

	if basepath == "/" {
		return "", errors.New("file not found")
	}

	// find up
	nearest := filepath.Dir(basepath)
	return lookup(nearest, filenames)
}
