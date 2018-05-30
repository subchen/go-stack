package filewalk

import (
	"os"
	"path/filepath"
)

// SkipDirFunc is the function called for each directory, return true to skip dir.
//
// The param path is an absolute path
type SkipDirFunc func(path string, info os.FileInfo) bool

// AcceptFunc is the function called for each file, return true to accpet file.
//
// The param path is an absolute patn
type AcceptFunc func(path string, info os.FileInfo) bool

// FindFiles finds matches file in root dir.
//
// if skipDirFn is nil, no dir waill be skipped.
// if relative is true, return relative filename.
func FindFiles(root string, acceptFn AcceptFunc, skipDirFn SkipDirFunc, relative bool) ([]string, error) {
	matches := make([]string, 0, 8)

	if !filepath.IsAbs(root) {
		root, _ = filepath.Abs(root)
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// dir
		if info.IsDir() {
			if skipDirFn != nil && skipDirFn(path, info) {
				return filepath.SkipDir
			}
			return nil
		}

		// file
		if acceptFn(path, info) {
			filename := path
			if relative {
				filename, _ = filepath.Rel(root, filename)
			}
			matches = append(matches, filename)
		}

		return nil
	})

	return matches, err
}
