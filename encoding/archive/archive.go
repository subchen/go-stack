// Package archive provides tar.gz and zip archiving
package archive

import (
	"path/filepath"

	"github.com/subchen/go-stack/archive/tar"
	"github.com/subchen/go-stack/archive/zip"
)

// Archive represents a compression archive files from disk can be written to.
type Archive interface {
	Add(name, path string) error
	Close() error
}

// New archive
// If the exentions of the target file is .zip, the archive will be in the zip
// format, otherwise, it will be a tar.gz archive.
func New(filename string) Archive {
	if filepath.Ext(filename) == ".zip" {
		return zip.New(filename)
	}
	return tar.New(filename)
}
